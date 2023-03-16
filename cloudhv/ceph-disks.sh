gcloud compute disks create ceph-0-disk-0 \
  --size 50 \
  --type "pd-ssd" \
  --zone australia-southeast2-a
gcloud compute disks create ceph-0-disk-1 \
  --size 50 \
  --type "pd-ssd" \
  --zone australia-southeast2-a
gcloud compute disks create ceph-0-disk-2 \
  --size 50 \
  --type "pd-ssd" \
  --zone australia-southeast2-a
gcloud compute instances attach-disk ceph-0 --disk ceph-0-disk-0
gcloud compute instances attach-disk ceph-0 --disk ceph-0-disk-1
gcloud compute instances attach-disk ceph-0 --disk ceph-0-disk-2
