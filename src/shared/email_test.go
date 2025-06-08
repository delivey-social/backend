package shared_test

import (
	"testing"

	"comida.app/src/shared"
)

var validEmails = []string{
	"thiagotolotti@gmail.com",
	"teste@teste.com",
	"alice@example.com",
	"bob.smith@domain.org",
	"carol.jones@company.co",
	"dave123@service.net",
	"eve_2024@sample.io",
	"frank+test@demo.com",
	"grace.hopper@navy.mil",
	"henry@university.edu",
	"irene@startup.ai",
	"jack_black@music.fm",
	"kate-winslet@movies.tv",
	"leo.dicaprio@hollywood.com",
	"maria.garcia@espanol.es",
	"nancy_drew@mystery.org",
	"oscar@awards.com",
	"paul.paulson@family.biz",
	"quincy@letters.xyz",
	"rachel.green@friends.tv",
	"samuel.jackson@actors.net",
	"tina.turner@music.com",
	"uma.thurman@cinema.org",
	"victor@winners.club",
	"wendy@fastfood.com",
	"xavier@school.edu",
	"yolanda@dance.studio",
	"zachary@alphabet.com",
	"anna.bell@bells.net",
	"brian.o'connor@cars.com",
	"claire@fashion.style",
	"daniel@coding.dev",
	"emily@flowers.shop",
	"felix@cats.org",
	"george@jungle.com",
	"harry.potter@hogwarts.edu",
	"isabel@travel.agency",
	"julia@recipes.cook",
	"kevin@homealone.com",
	"linda@wellness.health",
	"michael@basketball.net",
	"nina@art.gallery",
	"oliver@books.store",
	"peter.parker@spiderman.com",
	"queen@royalty.uk",
	"roger@tennis.pro",
	"sophia@language.school",
	"tom@cruise.com",
	"ursula@sea.org",
	"victoria@secrets.com",
	"username@yahoo.corporate",
	"username@domain.toolongtld",
}

var invalidEmails = []string{
	"",
	"plainaddress",
	"@missingusername.com",
	"username@.com",
	"username@com",
	"username@domain..com",
	"username@domain,com",
	"username@domain@domain.com",
	"username@domain",
	"username@.domain.com",
	".username@yahoo.com",
	"username@yahoo.com.",
	"username@yahoo..com",
	"username@yahoo.c",
	"username@-domain.com",
	"username@domain-.com",
	"username@domain_.com",
	"username@domain#.com",
	"username@domain..com",
	"username@.com.com",
	"username@domain..com",
	"username@domain..co.uk",
	"username@.domain.com",
	"username@domain.com.",
	"username@domain..com",
	"username@domain,com",
	"username@domain@domain.com",
	"username@domain",
	"username@.domain.com",
	"username@domain..com",
	"username@domain.c_m",
	"username@domain.c",
	"username@domain..",
	"username@.domain..com",
	"username@domain..com.",
	"username@.domain.com.",
	"username@domain..com..",
	"username@.com",
	"username@domain.com..",
	"username@.domain.com..",
	"username@domain..com.",
	"username@domain..com..",
	"username@.domain.com..",
	"username@domain..com.",
	"username@domain..com..",
	"username@.domain.com..",
	"username@domain..com.",
	"username@domain..com..",
	"username@.domain.com..",
}

func TestEmail(t *testing.T) {
	t.Run("Valid emails", func(t *testing.T) {
		for _, input := range validEmails {
			_, err := shared.NewEmail(input)
			if err != nil {
				t.Errorf("Unexpected error for %s: %v", input, err)
			}
		}
	})

	t.Run("Invalid emails", func(t *testing.T) {
		for _, input := range invalidEmails {
			_, err := shared.NewEmail(input)
			if err == nil {
				t.Errorf("Expected error for %s got nothing", input)
			}
		}
	})
}
