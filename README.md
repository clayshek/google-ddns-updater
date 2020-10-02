# google-ddns-updater
GoLang Google Domains Dynamic DNS Update Client

## Summary

google-ddns-updater is a program, writen in Go, for updating Google Domains dynamic DNS records via the Google Domains API. See Google documentation for API usage and other details: <a href="https://support.google.com/domains/answer/6147083?hl=en">https://support.google.com/domains/answer/6147083?hl=en</a>

ggoogle-ddns-updater will compare the current external IP address as obtained from http://myexternalip.com, with the currently resolving IP address of the DNS record. If they are identical, no further action is taken other than logging. If they differ, an authenticated request is sent to the Google Domains API to update the record, with logging of the request response and status.

## Requirements

- An existing dynamic DNS record hosted by Google Domains
- Environment variables set for GOOG_DDNS_HOSTNAME, GOOG_DDNS_USERNAME, & GOOG_DDNS_PASSWORD

## Usage

### Standalone Binary 
- Clone this repository
- Build the binary from src for your architecture:  <code>go build</code>
- Set environment variables for GOOG_DDNS_HOSTNAME, GOOG_DDNS_USERNAME, & GOOG_DDNS_PASSWORD
- Run google-ddns-updater

### Container Image
- Updates in this repo are automatically built via Dockerfile and a GitHub Action into a container image (ARM architecture) located at https://github.com/users/clayshek/packages/container/package/google-ddns-updater. <code> docker pull ghcr.io/clayshek/google-ddns-updater:latest </code>. Required environment variables will need to be provided to the container.

### Kubernetes
- The included <code>ddns-k8s-deployment.yml</code> file can be used to pull the continer image into a Kubernetes or K3s cluster, which runs as an hourly <a href="https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/">CronJob</a>. The secrets must be populated in the deployment file to be passed as environment variables to the pod.

## To-Do

 - [ ] Improve logging
 - [ ] Add better code comments
 - [ ] Create Helm chart
