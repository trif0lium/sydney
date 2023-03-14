gcloud compute instances create wynyard-0 \
  --enable-nested-virtualization \
  --zone australia-southeast2-a \
  --machine-type "n1-standard-32" \
  --boot-disk-size "100GB" \
  --image-family "debian-11"
