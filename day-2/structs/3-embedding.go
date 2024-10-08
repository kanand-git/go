package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) UpdateEmail(email string) {
	u.email = email
}

type bookAuthor struct {
	// embedding user struct in bookAuthor
	user  // anonymous field // a field with no name
	email string
	bio   string
}

type movieDirector struct {
	u         user // u is field name, user is type, not embedding
	movieName string
}

func (b *bookAuthor) UpdateBio(s string) {
	b.bio = s
}

func main() {
	author := bookAuthor{
		user: user{
			name:  "John",
			email: "john@email.com",
		},
		bio: "A book author",
	}

	// accessing the UpdateEmail of the user struct directly, as it is embedded in the bookAuthor struct
	author.UpdateEmail("john@gmail.com")
	author.UpdateBio("A book author with a new bio")
	fmt.Println(author)

	m := movieDirector{
		u: user{
			name:  "Director",
			email: "director@email.com",
		},
		movieName: "famous movie",
	}

	// accessing the UpdateEmail using u variable of the user struct
	m.u.UpdateEmail("director@gmail.com")

}
