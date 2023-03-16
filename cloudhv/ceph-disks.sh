gcloud compute instances create ceph-0 \
  --enable-nested-virtualization \
  --zone australia-southeast2-a \
  --machine-type "n1-standard-32" \
  --boot-disk-size "500GB" \
  --image-family "debian-11" \
  --image-project "debian-cloud"
