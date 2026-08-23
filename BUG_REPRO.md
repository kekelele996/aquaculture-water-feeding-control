# Bug Reproduction

## Symptom

Creating a quarantine pond and then adding its field metadata can terminate the
request before a pond is saved.

## Trigger

1. Create a pond from the quarantine intake flow without optional metadata.
2. Add a field value during the same request.
3. Repeat the flow with a policy value supplied through an interface.

## Observed Error

```
panic: runtime error: assignment to entry in nil map
```
