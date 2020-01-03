package dynamodb

type DynamoDbClient struct {
	tableName string
}

func New(tableName string) *DynamoDbClient {
	return &DynamoDbClient{
		tableName: tableName,
	}
}

func (d *DynamoDbClient) Save() string {
	return d.tableName
}

func (d DynamoDbClient) Find() string {
	return "find"
}

func (d DynamoDbClient) Query() string {
	return "query"
}

func (d DynamoDbClient) Scan() string {
	return "scan"
}

func (d DynamoDbClient) Update() string {
	return "updated"
}

func (d DynamoDbClient) UpdateItem() string {
	return "updated"
}

func (d DynamoDbClient) RemoveItem() string {
	return "removed"
}
