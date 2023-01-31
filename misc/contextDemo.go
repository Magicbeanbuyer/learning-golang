package misc

// https://www.educative.io/answers/what-are-tags-in-golang
import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = enrichContext(ctx)
	fetchContext(ctx)
}

// best practice: put context as the first parameter
func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "traceId", "123456")
}

func fetchContext(ctx context.Context) {
	t := ctx.Value("traceId")
	fmt.Printf("The trace ID is %v\n", t)
}
