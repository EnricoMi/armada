package utilisation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	armadaresource "github.com/armadaproject/armada/internal/common/resource"
	"github.com/armadaproject/armada/internal/common/util"
	"github.com/armadaproject/armada/internal/executor/domain"
	"github.com/armadaproject/armada/pkg/api"
	"github.com/armadaproject/armada/pkg/executorapi"
)

const nodeIdLabel = "node-id"

func TestGetAllPodsUsingResourceOnProcessingNodes_ShouldExcludePodsNotOnGivenNodes(t *testing.T) {
	presentNodeName := "Node1"
	podOnNode := v1.Pod{
		Spec: v1.PodSpec{
			NodeName: presentNodeName,
		},
	}
	podNotOnNode := v1.Pod{
		Spec: v1.PodSpec{
			NodeName: "Node2",
		},
	}

	node := v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: presentNodeName,
		},
	}
	pods := []*v1.Pod{&podOnNode, &podNotOnNode}
	nodes := []*v1.Node{&node}

	result := getAllPodsRequiringResourceOnNodes(pods, nodes)

	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0].Spec.NodeName, presentNodeName)
}

func TestGetAllPodsUsingResourceOnProcessingNodes_ShouldHandleNoNodesProvided(t *testing.T) {
	podOnNode := v1.Pod{
		Spec: v1.PodSpec{
			NodeName: "Node1",
		},
	}

	pods := []*v1.Pod{&podOnNode}
	var nodes []*v1.Node

	result := getAllPodsRequiringResourceOnNodes(pods, nodes)

	assert.Equal(t, len(result), 0)
}

func TestGetAllPodsUsingResourceOnProcessingNodes_ShouldIncludeManagedPodsOnNodes(t *testing.T) {
	podOnNode := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{domain.JobId: "label"},
		},
		Spec: v1.PodSpec{
			NodeName: "Node1",
		},
	}
	node := v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "Node1",
		},
	}

	pods := []*v1.Pod{&podOnNode}
	nodes := []*v1.Node{&node}

	result := getAllPodsRequiringResourceOnNodes(pods, nodes)

	assert.Equal(t, len(result), 1)
}

func TestGetAllPodsUsingResourceOnProcessingNodes_ShouldIncludeManagedPodNotAssignedToAnyNode(t *testing.T) {
	podOnNode := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{domain.JobId: "label"},
		},
		Spec: v1.PodSpec{
			NodeName: "",
		},
	}
	node := v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "Node1",
		},
	}

	pods := []*v1.Pod{&podOnNode}
	nodes := []*v1.Node{&node}

	result := getAllPodsRequiringResourceOnNodes(pods, nodes)

	assert.Equal(t, len(result), 1)
}

func TestGetAllPodsUsingResourceOnProcessingNodes_ShouldExcludeManagedPodNotAssignedToGivenNodes(t *testing.T) {
	podOnNode := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{domain.JobId: "label"},
		},
		Spec: v1.PodSpec{
			NodeName: "Node2",
		},
	}
	node := v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "Node1",
		},
	}

	pods := []*v1.Pod{&podOnNode}
	nodes := []*v1.Node{&node}

	result := getAllPodsRequiringResourceOnNodes(pods, nodes)

	assert.Equal(t, len(result), 0)
}

func TestGetUsageByQueue_HasAnEntryPerQueue(t *testing.T) {
	var priority int32
	podResource := makeResourceList(2, 50)
	queue1Pod1 := makePodWithResource("queue1", podResource, &priority)
	queue1Pod2 := makePodWithResource("queue1", podResource, &priority)
	queue2Pod1 := makePodWithResource("queue2", podResource, &priority)

	pods := []*v1.Pod{&queue1Pod1, &queue1Pod2, &queue2Pod1}

	result := GetAllocationByQueue(pods)
	assert.Equal(t, len(result), 2)
	assert.True(t, hasKey(result, "queue1"))
	assert.True(t, hasKey(result, "queue2"))
}

func TestGetUsageByQueue_SkipsPodsWithoutQueue(t *testing.T) {
	var priority int32
	podResource := makeResourceList(2, 50)
	pod := makePodWithResource("", podResource, &priority)
	pod.Labels = make(map[string]string)

	result := GetAllocationByQueue([]*v1.Pod{&pod})
	assert.NotNil(t, result)
	assert.Equal(t, len(result), 0)
}

func TestGetUsageByQueue_AggregatesPodResourcesInAQueue(t *testing.T) {
	var priority int32
	podResource := makeResourceList(2, 50)
	queue1Pod1 := makePodWithResource("queue1", podResource, &priority)
	queue1Pod2 := makePodWithResource("queue1", podResource, &priority)

	pods := []*v1.Pod{&queue1Pod1, &queue1Pod2}

	expectedResource := makeResourceList(4, 100)
	expectedResult := map[string]armadaresource.ComputeResources{"queue1": armadaresource.FromResourceList(expectedResource)}

	result := GetAllocationByQueue(pods)
	assert.Equal(t, result, expectedResult)
}

func TestGetUsageByQueue_HandlesEmptyList(t *testing.T) {
	var pods []*v1.Pod

	result := GetAllocationByQueue(pods)

	assert.NotNil(t, result)
	assert.Equal(t, len(result), 0)
}

func TestGetRunIdsByNode(t *testing.T) {
	utilisationService := &ClusterUtilisationService{
		nodeIdLabel: nodeIdLabel,
	}
	node1 := &v1.Node{ObjectMeta: metav1.ObjectMeta{Name: "node-1", Labels: map[string]string{nodeIdLabel: "node-1-id"}}}
	node2 := &v1.Node{ObjectMeta: metav1.ObjectMeta{Name: "node-2", Labels: map[string]string{nodeIdLabel: "node-2-id"}}}

	tests := map[string]struct {
		inputPods      []*v1.Pod
		expectedOutput map[string]map[string]api.JobState
	}{
		"MatchesOnNodeName": {
			inputPods: []*v1.Pod{
				createPodOnNode("job-1", "run-1", v1.PodRunning, "node-1", ""),
				createPodOnNode("job-2", "run-2", v1.PodRunning, "node-1", ""),
				createPodOnNode("job-3", "run-3", v1.PodRunning, "node-2", ""),
			},
			expectedOutput: map[string]map[string]api.JobState{
				"node-1": {"run-1": api.JobState_RUNNING, "run-2": api.JobState_RUNNING},
				"node-2": {"run-3": api.JobState_RUNNING},
			},
		},
		"HandlesAllPodPhases": {
			inputPods: []*v1.Pod{
				createPodOnNode("job-1", "run-1", v1.PodPending, "node-1", ""),
				createPodOnNode("job-2", "run-2", v1.PodRunning, "node-1", ""),
				createPodOnNode("job-3", "run-3", v1.PodSucceeded, "node-1", ""),
				createPodOnNode("job-4", "run-4", v1.PodFailed, "node-1", ""),
			},
			expectedOutput: map[string]map[string]api.JobState{
				"node-1": {"run-1": api.JobState_PENDING, "run-2": api.JobState_RUNNING, "run-3": api.JobState_SUCCEEDED, "run-4": api.JobState_FAILED},
			},
		},
		"PodWithNodeSelectorTargetingNode": {
			inputPods: []*v1.Pod{
				// Node selector matches node-1 label
				createPodOnNode("job-1", "run-1", v1.PodPending, "", "node-1-id"),
			},
			expectedOutput: map[string]map[string]api.JobState{
				"node-1": {"run-1": api.JobState_PENDING},
			},
		},
		"PodWithNodeSelectorTargetingInvalidNode": {
			inputPods: []*v1.Pod{
				// Node selector does not match any node
				createPodOnNode("job-1", "run-1", v1.PodPending, "", "node-3-id"),
			},
			// No matches
			expectedOutput: map[string]map[string]api.JobState{},
		},
		"Mixed": {
			inputPods: []*v1.Pod{
				createPodOnNode("job-1", "run-1", v1.PodRunning, "node-1", ""),
				createPodOnNode("job-2", "run-2", v1.PodPending, "", "node-1-id"),
			},
			expectedOutput: map[string]map[string]api.JobState{
				"node-1": {"run-1": api.JobState_RUNNING, "run-2": api.JobState_PENDING},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := utilisationService.getRunIdsByNode([]*v1.Node{node1, node2}, tc.inputPods)
			assert.Equal(t, tc.expectedOutput, result)
		})
	}
}

func createPodOnNode(jobId string, runId string, phase v1.PodPhase, nodeName string, nodeIdSelector string) *v1.Pod {
	pod := &v1.Pod{
		Status: v1.PodStatus{
			Phase: phase,
		},
		Spec: v1.PodSpec{
			NodeName: nodeName,
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				domain.JobId:    jobId,
				domain.JobRunId: runId,
			},
		},
	}
	if nodeIdSelector != "" {
		pod.Spec.NodeSelector = map[string]string{nodeIdLabel: nodeIdSelector}
	}
	return pod
}

func TestGetAllocatedResourceByNodeName(t *testing.T) {
	var priority int32
	podResource := makeResourceList(2, 50)
	pod1 := makePodWithResource("queue1", podResource, &priority)
	pod2 := makePodWithResource("queue1", podResource, &priority)
	pod3 := makePodWithResource("queue1", podResource, &priority)
	pod1.Spec.NodeName = "node1"
	pod2.Spec.NodeName = "node2"
	pod3.Spec.NodeName = "node2"
	pods := []*v1.Pod{&pod1, &pod2, &pod3}

	allocatedResource := getAllocatedResourceByNodeName(pods)
	assert.Equal(t, map[string]armadaresource.ComputeResources{
		"node1": armadaresource.FromResourceList(makeResourceList(2, 50)),
		"node2": armadaresource.FromResourceList(makeResourceList(4, 100)),
	}, allocatedResource)
}

func hasKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

func makeResourceList(cores int64, gigabytesRam int64) v1.ResourceList {
	cpuResource := resource.NewQuantity(cores, resource.DecimalSI)
	memoryResource := resource.NewQuantity(gigabytesRam*1024*1024*1024, resource.DecimalSI)
	resourceMap := map[v1.ResourceName]resource.Quantity{
		v1.ResourceCPU:    *cpuResource,
		v1.ResourceMemory: *memoryResource,
	}
	return resourceMap
}

func makePodWithResource(queue string, resource v1.ResourceList, priority *int32) v1.Pod {
	pod := v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Resources: v1.ResourceRequirements{
						Requests: resource,
						Limits:   resource,
					},
				},
			},
			Priority: priority,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: util.NewULID(),
		},
		Status: v1.PodStatus{
			Phase: v1.PodRunning,
		},
	}
	if queue != "" {
		pod.ObjectMeta.Labels = map[string]string{domain.JobId: util.NewULID(), domain.Queue: queue}
	}
	return pod
}

func TestGetCordonedResource(t *testing.T) {
	nodes := []*v1.Node{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "node1",
			},
			Spec: v1.NodeSpec{
				Unschedulable: true,
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "node2",
			},
			Spec: v1.NodeSpec{
				Unschedulable: false,
			},
		},
	}
	pod1 := createTestPod("node1", "1", "10Gi")
	pod2 := createTestPod("node1", "2", "2Gi")
	pod3 := createTestPod("node2", "1", "10Gi")
	pods := []*v1.Pod{pod1, pod2, pod3}

	resources := getCordonedResource(nodes, pods)

	expected := armadaresource.ComputeResources{
		"cpu":    resource.MustParse("3"),
		"memory": resource.MustParse("12Gi"),
	}

	assert.True(t, expected.Equal(resources))
}

func TestCalculateNonArmadaResource(t *testing.T) {
	oneCpu := resource.MustParse("1")
	twoCpu := resource.MustParse("2")
	threeCpu := resource.MustParse("3")
	oneGi := resource.MustParse("1Gi")
	twoGi := resource.MustParse("2Gi")
	threeGi := resource.MustParse("3Gi")

	defaultMinToAllocatePriority := int32(3)
	defaultMinToAllocate := armadaresource.ComputeResources{
		"cpu":    twoCpu,
		"memory": twoGi,
	}

	tests := map[string]struct {
		nonArmadaPods         []*v1.Pod
		minToAllocate         armadaresource.ComputeResources
		minToAllocatePriority int32
		expected              map[int32]*executorapi.ComputeResource
	}{
		"no pods": {
			nonArmadaPods:         []*v1.Pod{},
			minToAllocate:         defaultMinToAllocate,
			minToAllocatePriority: defaultMinToAllocatePriority,
			expected: map[int32]*executorapi.ComputeResource{
				defaultMinToAllocatePriority: {
					Resources: map[string]*resource.Quantity{
						"cpu":    &twoCpu,
						"memory": &twoGi,
					},
				},
			},
		},
		"One pod below minimum": {
			nonArmadaPods:         []*v1.Pod{createTestPod("", "1", "1Gi")},
			minToAllocate:         defaultMinToAllocate,
			minToAllocatePriority: defaultMinToAllocatePriority,
			expected: map[int32]*executorapi.ComputeResource{
				0: {
					Resources: map[string]*resource.Quantity{
						"cpu":    &oneCpu,
						"memory": &oneGi,
					},
				},
				defaultMinToAllocatePriority: {
					Resources: map[string]*resource.Quantity{
						"cpu":    &oneCpu,
						"memory": &oneGi,
					},
				},
			},
		},
		"One pod above minimum": {
			nonArmadaPods:         []*v1.Pod{createTestPod("", "3", "3Gi")},
			minToAllocate:         defaultMinToAllocate,
			minToAllocatePriority: defaultMinToAllocatePriority,
			expected: map[int32]*executorapi.ComputeResource{
				0: {
					Resources: map[string]*resource.Quantity{
						"cpu":    &threeCpu,
						"memory": &threeGi,
					},
				},
			},
		},
		"two pods same priority": {
			nonArmadaPods: []*v1.Pod{
				createTestPod("", "1", "1Gi"),
				createTestPod("", "2", "2Gi"),
			},
			minToAllocate:         defaultMinToAllocate,
			minToAllocatePriority: defaultMinToAllocatePriority,
			expected: map[int32]*executorapi.ComputeResource{
				0: {
					Resources: map[string]*resource.Quantity{
						"cpu":    &threeCpu,
						"memory": &threeGi,
					},
				},
			},
		},
		"two pods different priority": {
			nonArmadaPods: []*v1.Pod{
				createTestPod("", "1", "1Gi"),
				createTestPodAtPriority("", "1", "1Gi", defaultMinToAllocatePriority),
			},
			minToAllocate:         defaultMinToAllocate,
			minToAllocatePriority: defaultMinToAllocatePriority,
			expected: map[int32]*executorapi.ComputeResource{
				0: {
					Resources: map[string]*resource.Quantity{
						"cpu":    &oneCpu,
						"memory": &oneGi,
					},
				},
				defaultMinToAllocatePriority: {
					Resources: map[string]*resource.Quantity{
						"cpu":    &oneCpu,
						"memory": &oneGi,
					},
				},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := calculateNonArmadaResource(tc.nonArmadaPods, tc.minToAllocate, tc.minToAllocatePriority)

			require.Equal(t, len(tc.expected), len(result), "uneven priorities")
			for priority, expectedCr := range tc.expected {
				actualCr, ok := result[priority]
				require.True(t, ok, "No resources found at priority %d", priority)
				actualResources := actualCr.Resources
				require.Equal(t, len(expectedCr.Resources), len(actualResources), "uneven resources at priority %d", priority)
				for resName, expectedQty := range expectedCr.Resources {
					actualQty, ok := actualResources[resName]
					require.True(t, ok, "Resource %s not found at priority %d", resName, priority)
					assert.Equal(t, expectedQty.MilliValue(), actualQty.MilliValue(), "expected %s qty [%d] differs from actual [%d] at priority %d, ", resName, expectedQty.MilliValue(), actualQty.MilliValue(), priority)
				}
			}
		})
	}
}

func createTestPod(nodeName string, cpu string, memory string) *v1.Pod {
	return &v1.Pod{
		Spec: v1.PodSpec{
			NodeName: nodeName,
			Containers: []v1.Container{
				{
					Resources: v1.ResourceRequirements{
						Requests: v1.ResourceList{
							"cpu":    resource.MustParse(cpu),
							"memory": resource.MustParse(memory),
						},
						Limits: v1.ResourceList{
							"cpu":    resource.MustParse(cpu),
							"memory": resource.MustParse(memory),
						},
					},
				},
			},
		},
	}
}

func createTestPodAtPriority(nodeName string, cpu string, memory string, priority int32) *v1.Pod {
	pod := createTestPod(nodeName, cpu, memory)
	pod.Spec.Priority = &priority
	return pod
}
