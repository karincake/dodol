# Dodol - Data Structure
Collection of data structre to organize ouput.

DISCLAIMER: This is a very opinionated package

## Data Type
- `IS` map `string` keyed of `string`
- `ISIS` map `string` keyed of `IS`
- `II` map string keyed of `interface`
- `Data` structured type intended for data in general containing `Meta`, `Data`, and `Ref` (for reference)
- `Message` structured message type data containing `Code` and `Message`
- `Content` structured data for content only.

## Common Error Type
- `CommonError` structured type intended for `error` containing `code` and `Message`
- `CommonErrors` map `string` keyed of `CommonError`

## Field Error
- `FieldError` structured type intended for field `error` containing `code`, `Message`, `ExpectedVal`, `GivenVal`, and `EmbedSource`
- `FieldErrors` map `string` keyed of `FieldError`
