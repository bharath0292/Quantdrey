# 📝 TODO

## Global Improvements
- [ ] **Graceful Shutdowns**
  - Implement graceful shutdowns for the application.
  - Ensure all resources are properly released and connections are closed.

- [ ] **Remove Unwanted, Unnecessary Context**
  - Remove any unnecessary context from the codebase.
  - Ensure that context is only used where necessary.

## Strategy Domain Improvements
- [ ] **Custom Validation for Strategy Domain (CRU)**
  - Add validation rules for all `Strategy` fields.
  - Ensure nested entities are validated, including:
    - `Rule`
    - `Symbol`
    - `EntryLogic`
    - `ExitLogic`
  - Apply consistently for **Create**, **Read**, and **Update** operations.
  - Use custom GraphQL scalars or service‑layer validation as needed.

- [ ] **Nullable Field Handling in Strategy Update**
  - **Current behavior:**
    - If a user sends `null` for a field, the update skips that field and the existing value remains.
  - **Required behavior:**
    - **Explicit `null`** → clear/remove the field in DB (`$unset` in MongoDB).
    - **Omitted field** → leave the value unchanged.
  - **Implementation hint:**
    - Enhance the recursive update builder:
      - `$set` → for non‑nil values.
      - `$unset` → for present-but-`nil` pointers.
