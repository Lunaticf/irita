<!--
order: false
-->

### 资产元数据 Schema 示例

```json
{
    "title": "Asset Metadata",
    "type": "object",
    "properties": {
        "name": {
            "type": "string",
            "description": "Identifies the asset to which this NFT represents"
        },
        "description": {
            "type": "string",
            "description": "Describes the asset to which this NFT represents"
        },
        "image": {
            "type": "string",
            "description": "A URI pointing to a resource with mime type image/* 
                representing the asset to which this NFT represents"
        },
        "meta": {
            "type": "string",
            "description": "A URI pointing to other attributes related to the asset to which this NFT represents"
        }
    }
}
```
