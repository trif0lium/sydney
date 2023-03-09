gcloud compute instances create cloudhv-dev-vm \
  --enable-nested-virtualization \
  --zone australia-southeast2-a \
  --machine-type "n2-standard-8" \
  --boot-disk-size "100GB"
