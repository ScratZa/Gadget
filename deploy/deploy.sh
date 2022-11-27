kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default
kubectl apply -f deploy/gulper.yaml
