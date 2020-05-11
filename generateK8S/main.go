package main

/**
This is a sample of how the pods can be created given some environment variables
This way is simpler than replacing the variables in a templatized yaml since this
leaves some room for extending the project and allows other developer flexibility
in terms of version changes in k8s APIs.
*/
import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
)

func main() {
	pod := createPod("dev")
	bytes, err := json.Marshal(pod)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func createPod(environment string) *corev1.Pod {
	return &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-pod",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "test-container",
					Image: "nginx",
					Env: []corev1.EnvVar{
						{
							Name:  "ENV",
							Value: environment,
						},
					},
				},
			},
		},
	}
}
