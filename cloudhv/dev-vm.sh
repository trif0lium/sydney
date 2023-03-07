gcloud compute instances create cloudhv-dev-vm \
  --enable-nested-virtualization \
  --zone asia-southeast1-a \
  --min-cpu-platform "Intel Haswell" \
  --machine-type "n2-standard-4"
