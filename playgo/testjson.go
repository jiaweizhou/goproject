package main

import (
	"K8APITransform/K8APITransform/ApiServer/models"
	"encoding/json"
	"fmt"
)

func main() {
	json1 := `{
        "replicas": 1,
        "selector": {
          "version": "test-1"
        },
        "template": {
          "metadata": {
            "creationTimestamp": null,
            "labels": {
              "name": "test",
              "version": "test-1"
            }
          },
          "spec": {
            "volumes": null,
            "containers": [
              {
                "name": "php-red",
                "image": "kubernetes/example-guestbook-php-redis:v2",
                "ports": [
                  {
                    "containerPort": 80,
                    "protocol": "TCP"
                  }
                ],
                "resources": {},
                "terminationMessagePath": "/dev/termination-log",
                "imagePullPolicy": "IfNotPresent",
                "capabilities": {}
              }
            ],
            "restartPolicy": "Always",
            "dnsPolicy": "ClusterFirst"
          }
        }
      }`
	var b models.ReplicationControllerSpec
	json.Unmarshal([]byte(json1), &b)
	fmt.Println(b.Template.ObjectMeta.Labels)
}
