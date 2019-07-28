# Overview
This is a WebService that provides access to the UK HMRC MTD(MakeTaxDigital) APIs. 
The source code for this shim API is hosted under https://hmrc_oauth.awltux.trader.
The intention is to allow Open Source Accounting tools to call 
this shim API and provide access to the Open Source unfriendly MTD API.

# MTD API
The MTD API is currently used to submit VAT returns from VAT registered organisations.
The MTD API is protected by two mechanisms that make it difficult to use with
open source accounting products:
1. Developer ID: All calls to the API must contain a valid, pre-registered Devloper ID
2. Callback REST API: This is a pre-registered web address to a SSL secured REST API.

## Developer ID
A developer ID can be created by anyone registering with the HMRC Developer 