---
name: Bug report
about: Create a report to help us improve
title: ''
labels: bug
assignees: wprins-JNPR

---

## Describe the bug
A clear and concise description of what the bug is.


## To Reproduce
Steps to reproduce the behavior:
1. Create, update or refresh '.....' resource or datasource
2. Configure the resource/datasource with the following HCL configuration (paste the HCL configuration, remove any sensitive information)
3. Use Terraform Apply / Refresh / ... command
4. See error


## Terraform configuration
An extract of the Terraform configuration related to this issue
```terraform
resource "mist_org" "terraform_test" {
  name = "Terraform Testing"
}
...
```

## Expected behavior
A clear and concise description of what you expected to happen.


## Error Message
If applicable, add screenshots or paste the Terraform error message to help explain your problem.


## Additional context
Add any other context about the problem here.
