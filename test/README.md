go-onedrive Tests
===============

This directory contains additional test suites beyond the unit tests already in
[../onedrive](../onedrive). Whereas the unit tests run very quickly (since they
don't make any network calls) and are run by GitHub Action on every commit, the tests
in this directory are only run manually.

The test packages are:

Integration
-----------

This will exercise the entire go-onedrive library (or at least as much as is
practical) against the live Microsoft Graph (and OneDrive API if Monitor is tested). 
These tests will verify that the library is properly coded against the actual behavior 
of the API, and will (hopefully) fail upon any incompatible change in the API.

Because these tests are running using live data, there is a much higher
probability of false positives in test failures due to network issues, test
data having been changed, etc.

These tests send real network traffic to the Microsoft Graph.
Additionally, in order to test the methods that modify data, a real OAuth token
will need to be present. While the tests will try to be well-behaved in terms
of what data they modify, it is **strongly** recommended that these tests only
be run using a dedicated test account.

Run the tests under the integration folder using:

```bash
SET MICROSOFT_GRAPH_ACCESS_TOKEN=XXX 
go test -v
```

Some of the tests are commented out because those test will create/delete/update the 
drive items and folders on the actual OneDrive. So, you can decide to uncomment them
when you are going to test those operations.

There is a keyword `<<input>>` in the test where it is for you to key in the actual value
based on your OneDrive setup, for example the actual Drive ID of your OneDrive Music 
folder.