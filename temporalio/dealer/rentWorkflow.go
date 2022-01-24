package main

import (
	"car-api/internal/core/domain"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"log"
	"time"
)

// workflow que implementa el patrón SAGA
// retorna la data de renta del auto
func rent(ctx workflow.Context, carRent *domain.CarRent) (*domain.CarRent, error) {

	// definimos un logger para el workflow
	logger := workflow.GetLogger(ctx)

	// definimos una opción de actividad general (por facilidad y simplicidad)
	// las opciones de actividad permiten configurar granularmente las actividades desde el punto de vista de Temporal
	// timeouts cortos en duración por propósitos del ejercicio
	aa := workflow.ActivityOptions{
		// tiempo máximo que puede transcurrir desde que un workflow solicite la ejecución de una actividad
		// hasta que un worker inicie la ejecución de dicha actividad.
		// Si se dispara este timeout es indicativo de que el/los workers que registran la actividad estan:
		// o abajo o no pueden mantener la velocidad de despacho de tareas.
		ScheduleToStartTimeout: time.Second * 20,

		// tiempo máximo dentro del cual se puede ejecutar una tarea una vez que es tomada por un worker.
		StartToCloseTimeout: time.Second * 20,
	}

	// inicializamos el contexto genérico y cargamos la opción de actividad general
	ctx = workflow.WithActivityOptions(ctx, aa)

	// INICIO DE IMPLEMENTACIÓN DE SAGA ----------------------------------------

	var carRented domain.CarRent
	var ctxCar = workflow.WithTaskQueue(ctx, "car") // indicamos que se enrute a la cola car

	errCar := workflow.ExecuteActivity(ctxCar, "car.rent", carRent).Get(ctxCar, &carRented) // indicamos que actividad ejecutar

	if errCar != nil {
		logger.Error("Falla ejecutando la actividad car.rent ", "Error", errCar)
		return nil, errCar
	}

	// FIN DE IMPLEMENTACIÓN DE SAGA -------------------------------------------

	return &carRented, nil
}

func main() {
	// Para iniciar debes abrir un objeto client. Este es un objeto pesado
	// Solo debes crear uno por proceso.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("No es posible crear el cliente", err)
	}
	defer c.Close()

	// creo el worker
	w := worker.New(c, "dealer_rent", worker.Options{}) // TaskQueue= car

	ow1 := workflow.RegisterOptions{
		Name: "dealer.rent",
	}
	w.RegisterWorkflowWithOptions(rent, ow1)

	// Ejecutar worker.
	// Este es un proceso demonio.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("No es posible ejecutar worker", err)
	}

}
