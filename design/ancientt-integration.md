---
title: "Ancientt Integration"
---

## Feature Description

* Allow users to start and cancel ancientt runs via the data-control-center.
* Ability to retrieve the logs and results file of the ancientt run.
* Only one test can run at a time.

## Flow

1. User starts the network test run.
   1. Output format choices: CSV and Excel.
2. Start ancientt run using command exec.
   1. A lock is set till the ancientt run is completed.
3. When completed store logs in a temporary file, and make sure the results file has been created.
   1. Unset the lock and set the last run time accordingly.
