package acc

import (
	"fmt"
	"testing"
	"math/rand"

	"github.com/stretchr/testify/require"

	"github.com/xanzy/go-gitlab"
)


func TestGitlabTopic_basic(t *testing.T) {
	client, _ := setup(t)
	testID := rand.Int()
	testTopicName := fmt.Sprintf("test-topic-%d", testID)

	// we can successfully get all topics
	_, _, err := client.Topics.ListTopics(&gitlab.ListTopicsOptions{})
	require.NoError(t, err)

	// we can create a topic
	topic, _, err := client.Topics.CreateTopic(&gitlab.CreateTopicOptions{Name: gitlab.String(testTopicName), Description: gitlab.String("Acceptance Test Topic")}, nil)
	require.NoError(t, err)

	// FIXME: topic's cannot yet be deleted, ...
	// make sure that the topic is deleted afterwards
	// defer func(topicID int) {
	// 	_, err = client.Topics.DeleteTopic(topicID)
	// 	require.NoError(t, err)
	// }(topic.ID)

	require.Equal(t, testTopicName, topic.Name)
	require.Equal(t, "Acceptance Test Topic", topic.Description)

	// we can query the topic
	gotTopic, _ ,err := client.Topics.GetTopic(topic.ID, nil)
	require.NoError(t, err)
	require.Equal(t, topic, gotTopic)

	// we see the topic in the list of topics
	topics, _, err := client.Topics.ListTopics(&gitlab.ListTopicsOptions{Search: gitlab.String(testTopicName)})
	require.NoError(t, err)
	require.Contains(t, topics, topic)
}