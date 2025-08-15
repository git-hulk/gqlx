# GQLX

A builder for GraphQL queries, mutations, and subscriptions. It's designed to
make the building process of GraphQL operations more intuitive and less error-prone.

NOTICE: this library is still in development and may change in the future.

## How to use

### Query Example

```Go
query := builder.Query().Name("user").
    AddSelections(
        NewField("id").Alias("user_id"),
        NewField("name").AddArguments(
            FromValue("age", value.Int(30)),
            FromValue("var", value.Variable("var")),
            FromType("sex", "Sex", nil),
            FromType("status", "UserStatus", value.String("active")),
        ),
    )

// query user {
//    user_id: id
//    name(age: 30, var: $var, sex: Sex, status: UserStatus = "active")
// }
queryString := query.String()
```

### Mutation Example

```Go
mutation := builder.Mutation().Name("CreateReviewForEpisode").AddArguments(
    FromType("ep", "Episode", nil),
    FromType("review", "ReviewInput", nil),
).AddSelections(
    NewField("createReview").AddArguments(
        FromValue("episode", value.Variable("ep")),
        FromValue("review", value.Variable("review")),
    ).AddSelections(
        NewField("stars"),
        NewField("commentary"),
    ),
)

// mutation CreateReviewForEpisode(ep: Episode, review: ReviewInput) {
//   createReview(episode: $ep, review: $review) {
//     stars
//     commentary
//   }
// }
mutationString := mutation.String()
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.