gcloud compute instances create wynyard-0 \
  --enable-nested-virtualization \
  --zone australia-southeast2-a \
  --machine-type "n1-standard-32" \
  --boot-disk-size "100GB" \
  --image-family "debian-11" \
  --image-project "debian-cloud"
gcloud compute disks create wynyard-0-disk-0 \
  --size 50 \
  --type "pd-ssd" \
  --zone australia-southeast2-a
gcloud compute instances attach-disk wynyard-0 --disk wynyard-0-disk-0
