# Provider Model

External catalogs/stream sources are isolated from the Shell.

## Contract

A provider may expose bounded HTTP/JSON or isolated-process RPC for:
- catalog/search;
- metadata;
- stream descriptors;
- subtitle descriptors.

The model is conceptually compatible with catalog/meta/stream/subtitle separation used by addon ecosystems, but project contracts remain under our control.

## Security

Providers do not receive arbitrary Shell execution or general filesystem access. Network and storage permissions are explicit per provider.

## Distribution

Official distributions include only lawful/public/authorized integrations. User-added providers remain subject to sandbox and content policy.
