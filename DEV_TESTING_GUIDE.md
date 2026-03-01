# Provider Development Testing Guide

This guide explains how to test local changes to the Terraform provider before submitting a PR.

## Prerequisites

- Go 1.21 or later
- Terraform 1.0 or later
- Access to a Mist organization (for testing)

## Setup for Local Development

### 1. Configure Terraform to Use Local Provider Build

Create a Terraform CLI configuration file to override the provider source:

**Linux/macOS:**
```bash
cat > ~/.terraformrc << 'EOF'
provider_installation {
  dev_overrides {
    "registry.terraform.io/juniper/mist" = "/Users/YOUR_USERNAME/go/bin"
  }

  # For all other providers, install them directly as normal.
  direct {}
}
EOF
```

**Windows:**
Create `%APPDATA%/terraform.rc` with:
```hcl
provider_installation {
  dev_overrides {
    "registry.terraform.io/juniper/mist" = "C:\\Users\\YOUR_USERNAME\\go\\bin"
  }

  direct {}
}
```

> **Note:** Replace `YOUR_USERNAME` with your actual username, or use the full path to your Go bin directory.

### 2. Build and Install the Provider

From the provider repository root:

```bash
# Build and install to $GOPATH/bin
go install .

```

This compiles the provider and places the binary in your Go bin directory (typically `~/go/bin`).

### 3. Verify the Override is Active

When you run Terraform commands, you should see:

```
Warning: Provider development overrides are in effect

The following provider development overrides are set in the CLI configuration:
 - juniper/mist in /Users/YOUR_USERNAME/go/bin

The behavior may therefore not match any released version of the provider...
```

This confirms Terraform is using your local build.

## Testing Workflow

### 1. Make Code Changes

Edit the provider source code in the `internal/` directory.

### 2. Rebuild the Provider

```bash
go install .
```

### 3. Create a Test Configuration

Create a `.tf` file with your test resources:

```hcl
terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/juniper/mist"
    }
  }
}

provider "mist" {
  # Configure via environment variables:
  # export MIST_API_TOKEN="your-token"
  # export MIST_HOST="api.mistsys.com"  # optional
}

# Add your test resources here
resource "mist_device_switch" "test" {
  site_id   = "your-site-id"
  device_id = "your-device-id"
  name      = "test-switch"
  # ... other attributes
}
```

### 4. Initialize and Test

```bash
# Initialize (pulls in other providers if needed)
terraform init

# Check the plan
terraform plan

# Apply changes (if safe to do so)
terraform apply

# Clean up
terraform destroy
```

### 5. Check for Errors

```bash
# Enable debug logging to see detailed output
export TF_LOG=DEBUG
terraform plan 2>&1 | tee terraform.log

# Or trace level for even more detail
export TF_LOG=TRACE
terraform apply 2>&1 | tee terraform-apply.log

# Disable logging
unset TF_LOG
```

## Best Practices

### Use Environment Variables for Credentials

Never hardcode credentials in your test configurations:

```bash
# Set in your shell (not in version control)
export MIST_API_TOKEN="your-token-here"
export MIST_HOST="api.mistsys.com"
```

### Use .tfvars Files for Test Data

Create a `test.tfvars` file (add to `.gitignore`):

```hcl
org_id    = "your-org-id"
site_id   = "your-site-id"
device_id = "your-device-id"
```

Then reference it:

```bash
terraform plan -var-file=test.tfvars
```

### Clean Up Test Resources

Always destroy test resources when done:

```bash
terraform destroy -auto-approve
```

### Isolate Test Configurations

Keep test `.tf` files separate from production:

```bash
# Use a dedicated testing directory
mkdir -p test/manual
cd test/manual
# Create test configs here
```

## Reverting to Released Provider

When you're done testing and want to use the published provider version:

### 1. Remove the Override Configuration

```bash
# Linux/macOS
rm ~/.terraformrc

# Windows
del %APPDATA%\terraform.rc
```

### 2. Reinitialize Terraform

```bash
# Remove local state (if this was just for testing)
rm -rf .terraform .terraform.lock.hcl terraform.tfstate*

# Download the released provider
terraform init
```

### 3. Verify Released Version is Used

Run `terraform version` - you should no longer see the development override warning.

## Troubleshooting

### Provider Not Found

If Terraform can't find the provider after `go install`:

1. Check your Go bin directory:
   ```bash
   echo $GOPATH/bin
   ls -la ~/go/bin/terraform-provider-mist*
   ```

2. Verify the path in `~/.terraformrc` matches

3. Rebuild:
   ```bash
   go install .
   ```

### Changes Not Taking Effect

1. Ensure you rebuilt the provider: `go install .`
2. Delete Terraform's plugin cache:
   ```bash
   rm -rf .terraform/providers
   terraform init
   ```

### Build Errors

```bash
# Check Go version
go version

# Verify dependencies
go mod tidy
go mod verify

# Clean and rebuild
go clean
go install .
```

### State Incompatibility

If you see "state is incompatible with published releases":

- This is expected when using dev overrides
- Test changes are isolated to your test environment
- Revert to released provider before applying to production

## Testing Checklist

Before submitting a PR:

- [ ] Code builds without errors: `go install .`
- [ ] Code passes linting: `golangci-lint run` (if available)
- [ ] Manual testing completed with test configuration
- [ ] No regressions in existing functionality
- [ ] Test resources cleaned up: `terraform destroy`
- [ ] `.tfvars` and state files not committed
- [ ] Development override removed: `rm ~/.terraformrc`

## Additional Resources

- [Terraform Plugin Development](https://developer.hashicorp.com/terraform/plugin)
- [terraform-plugin-framework Documentation](https://developer.hashicorp.com/terraform/plugin/framework)
- [Provider Testing Best Practices](https://developer.hashicorp.com/terraform/plugin/testing)
