package sqs

import (
	"accumulation_consumer/internal/domain"
	usecase "accumulation_consumer/internal/usercase"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSHandler struct {
	SQSClient      *sqs.SQS
	httpClientCase *usecase.HttpClientCase
	QueueURL       string
	ServiceUrl     string
}

func NewSQSHandler(queueURL string, serviceURL string, httpClientCase *usecase.HttpClientCase) *SQSHandler {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"), // Cambia esto a tu región de AWS deseada
	})
	if err != nil {
		log.Fatal("Error creando la sesión:", err)
	}

	sqsClient := sqs.New(sess)

	return &SQSHandler{
		SQSClient:      sqsClient,
		httpClientCase: httpClientCase,
		QueueURL:       queueURL,
		ServiceUrl:     serviceURL,
	}
}

func (h *SQSHandler) CreatePoint() error {
	for {
		receiveMessageInput := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(h.QueueURL),
			MaxNumberOfMessages: aws.Int64(1),
			WaitTimeSeconds:     aws.Int64(20),
		}

		result, err := h.SQSClient.ReceiveMessage(receiveMessageInput)
		if err != nil {
			log.Println("Error recibiendo el mensaje:", err)
		}

		for _, message := range result.Messages {
			point := &domain.Point{}
			err := json.Unmarshal([]byte(*message.Body), point)
			if err != nil {
				log.Println("Error al deserializar el mensaje:", err)
			}

			// Agrega tu lógica de negocio aquí
			log.Printf("Recibido punto: %+v\n", point)

			// Elimina el mensaje de la cola
			deleteMessageInput := &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(h.QueueURL),
				ReceiptHandle: message.ReceiptHandle,
			}
			_, err = h.SQSClient.DeleteMessage(deleteMessageInput)
			if err != nil {
				log.Println("Error eliminando el mensaje:", err)
			}
		}

		time.Sleep(5 * time.Second)
	}
}
