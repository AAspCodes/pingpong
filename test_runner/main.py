from kubernetes import client, config
import yaml

CLIENT_DEP_PATH = 'client/deployment.yaml'


def load_dep(path: str)->dict:
    with open(path, 'r') as f:
        dep = yaml.safe_load(f)
    print(type(dep))
    return dep



def create_client_pods(dep, n:int):
    # Load kube config
    config.load_kube_config()

    v1 = client.CoreV1Api()

    for i in range(5):  # replace with your desired number of Pods
        pod_name = f'{dep["metadata"]["name"]}-{i}'

        # Create the Pod
        pod = client.V1Pod(
            metadata=client.V1ObjectMeta(name=pod_name),
            spec=client.V1PodSpec(
                containers=[
                    client.V1Container(
                        name=pod_name,
                        image=f"{dep['spec']['template']['spec']['containers'][0]['image']}",
                        ports=[client.V1ContainerPort(container_port=8080)],
                        env=[client.V1EnvVar(name='POD_NAME', value=pod_name)]
                    )
                ]
            )
        )

        v1.create_namespaced_pod(namespace='default', body=pod)

        # # Create the Service
        # service = client.V1Service(
        #     metadata=client.V1ObjectMeta(name=f'{pod_name}-service-{i}'),
        #     spec=client.V1ServiceSpec(
        #         selector={'app': pod_name},
        #         ports=[client.V1ServicePort(port=80, target_port=8080)],
        #     )
        # )

        # v1.create_namespaced_service(namespace='default', body=service)

if __name__ == "__main__":
    dep = load_dep(CLIENT_DEP_PATH)
    create_client_pods(dep, 1)