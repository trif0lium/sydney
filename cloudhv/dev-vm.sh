gcloud compute instances create cloudhv-dev-vm \
  --enable-nested-virtualization \
  --zone australia-southeast2-a \
  --machine-type "n2-standard-8" \
  --boot-disk-size "100GB" \
  --image "projects/ubuntu-os-cloud/global/images/ubuntu-2210-kinetic-amd64-v20230125"
