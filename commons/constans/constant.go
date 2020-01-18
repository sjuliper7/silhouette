package constans

type Topic string

const (
	TopicUserRegistration Topic = "kafka.registration.finish"
	TopicUserUpdated      Topic = "user.updated.finish"
	TopicUserDeleted      Topic = "user.deleted.finish"
)
