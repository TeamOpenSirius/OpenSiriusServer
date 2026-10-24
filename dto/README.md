# Sirius API Go Structures

Generated from `dump.cs`.

Scope:
- `SiriusApi.Shared*`
- `SiriusRealtime.Shared*`
- Excludes `MessagePack.Formatters.*` and static utility classes without payload properties.

Generated counts:
- Structs/classes: 1001
- Enums: 188
- Properties: 4299
- Namespaces: 39

Notes:
- All files use package `siriusapi`.
- C# `[Key(n)]` values are kept as comments on fields and in `Schemas`.
- `json` and `msgpack` tags use lower-camel property names.
- Some dump-only helper types such as `Func<T,...>`, `RangeView<T>`, and unresolved interface types are represented as `any` so the package remains compilable.
- If the client expects MessagePack-CSharp integer-key array encoding, use `Schemas` to encode fields by numeric `Key` order instead of relying only on map-style struct tags.
- Duplicate C# type names across namespaces were renamed with namespace prefixes. See `Schemas` for original names.
