# try-connect

## Learnings

1. Running `buf lint` shows it has strict requirements on
   - Folder structure (where to put the protos)
   - Version number (e.g. "v1" and "v1alpha1" are ok but not "v0")
   - Service proto must be suffixed with "Service"
