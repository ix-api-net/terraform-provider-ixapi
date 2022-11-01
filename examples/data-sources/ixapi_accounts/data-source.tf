# Get all managed customer accounts
data "ixapi_accounts" "subcustomers" {
    managing_account = "<reseller_account_id>"
}

locals {
    subcustomers = data.ixapi_accounts.subcustomers.accounts
}

