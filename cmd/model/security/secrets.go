package security

type Secrets struct {
	SUPABASE_URL string
	SUPABASE_KEY string
	OPENAIKEY    string
}

var instance *Secrets

func GetInstance() *Secrets {

	if instance == nil {
		instance = &Secrets{}
	}

	return instance
}
