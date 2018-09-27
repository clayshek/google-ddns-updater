# google-ddns-updater
GoLang Google Domains Dynamic DNS Update Client

## Summary

google-ddns-updater is a client, writen in Go, for updating Google Domains dynamic DNS records via the Google Domains API. See Google documentation for API usage and other details: <a href="https://support.google.com/domains/answer/6147083?hl=en">https://support.google.com/domains/answer/6147083?hl=en</a>

ggoogle-ddns-updater will compare the current external IP address as obtained from http://myexternalip.com, with the currently resolving IP address of the DNS record. If they are identical, no further action is taken other than logging. If they differ, an authenticated request is sent to the Google Domains API to update the record, with logging of the request response and status.

## Requirements

- An existing dynamic DNS record hosted by Google Domains
- Environment variables set for GOOG_DDNS_HOSTNAME, GOOG_DDNS_USERNAME, & GOOG_DDNS_PASSWORD

## Usage

- Clone this repository
- Build the binary for your architecture:  <code>go build</code>
- Set environment variables for GOOG_DDNS_HOSTNAME, GOOG_DDNS_USERNAME, & GOOG_DDNS_PASSWORD
- Run google-ddns-updater


## To-Do

 - [ ] Improve metrics & logging?
 - [ ] Add better comments to code
