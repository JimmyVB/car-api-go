package main

import (
	"bufio"
	"car-api/internal/core/domain"
	"context"
	"encoding/json"
	"fmt"
	"go.temporal.io/sdk/client"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Para iniciar debes abrir un objeto client. Este es un objeto pesado
	// Solo debes crear uno por proceso.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("No es posible crear cliente", err)
	}
	defer c.Close()

	// Configura las opciones para el workflow
	// ID:
	// Definimos identificador para el workflow
	//
	// Definimos el nombre del TaskQueue.
	// Qué es TaskQueue?:
	// 		Cuando un workflow invoca una actividad,
	// 		se envía el comando ScheduleActivityTask al servicio de Temporal.
	// 		Como resultado, el servicio actualiza el estado del workflow y
	//		envía una tarea de actividad a un worker que implementa la actividad.
	//		En lugar de llamar al worker directamente, se utiliza una cola intermedia.
	// 		Entonces, el servicio agrega una tarea de actividad a esta cola y un worker
	// 		recibe la tarea mediante una solicitud de encuesta larga.
	// 		Temporal llama a esta cola que se utiliza para distribuir tareas de actividad
	// 		en una cola de tareas de actividad.
	workflowOptions := client.StartWorkflowOptions{
		ID:        "dealer_rent_workflowID", // identificador del workflow
		TaskQueue: "dealer_rent",            // enrutar a taskqueue del worker donde esta el workflow definido en este caso es dealer_rent.
	}

	// obtener datos para create auto
	strIdCar := StringPromptRent("IdCar: ")
	strIdUser := StringPromptRent("IdUser: ")
	strStartDate := StringPromptRent("StartDate: ")
	strEndDate := StringPromptRent("EndDate: ")
	idCar, _ := strconv.ParseInt(strIdCar, 0, 0)
	idUser, _ := strconv.ParseInt(strIdUser, 0, 0)
	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, strStartDate)
	endDate, _ := time.Parse(layout, strEndDate)
	carRent := domain.CarRent{IdCar: int(idCar), IdUser: int(idUser), StartDate: startDate, EndDate: endDate}
	// Ejecutamos el workflow.
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, "dealer.rent", &carRent)

	// si falla la ejecución
	if err != nil {
		log.Fatalln("No es posible ejecutar workflow", err)
	}

	// logueamos el Id del workflow y el Id de su ejecución
	log.Println("Workflow iniciado", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Para este caso esperamos sincronamente (también es posible asincrono) a que el workflow se complete.
	var result domain.CarRent
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("No es posible obtener resultado del workflow", err)
	}
	var finalMessage domain.Message
	message := "Se realizo la renta exitosamente"
	if result.IdCar == 0 {
		message = "No se ha podido realizar la renta del vehiculo, verifique los datos"
		finalMessage.Data = nil
	} else {
		finalMessage.Data = result
	}
	finalMessage.Message = message
	returnValue, _ := json.Marshal(finalMessage)
	log.Println("Resultado:", string(returnValue)) // info de renta
}

func StringPromptRent(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
