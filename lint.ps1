# Get root directory
$location = "$( Get-Location )"

# Run golangci-lint on the modified files.
golangci-lint run --out-format tab --path-prefix $location