package main

import (
	"context"
	"fmt"
	"github.com/loehnertz/Toranos/common"
	"github.com/loehnertz/Toranos/services/fleet-controller/proto"
	"github.com/micro/go-log"
	"time"
)

func checkForExpiredReservations() {
	resReservations, errReservations := fleetController.RetrieveReservations(context.TODO(), &fleet_controller.Empty{})

	if errReservations != nil {
		log.Log(errReservations)
	} else {
		for i := range resReservations.Reservations {
			reservation := resReservations.Reservations[i]
			createdAt := time.Unix(reservation.CreatedAt, 0)

			if reservation.Status == uint32(common.GetConfigIntByPath(conf, "service-settings", "fleet-controller", "statuses", "reserved")) &&
				time.Since(createdAt) > common.GetConfigDurationByPath(conf, "service-settings", "fleet-controller", "reservationTime") {
				resUnbook, errUnbook := fleetController.Unbook(context.TODO(), &fleet_controller.UnbookingRequest{
					VehicleId:  reservation.Vehicle,
					CustomerId: reservation.Customer,
				})
				if errUnbook != nil {
					log.Log(errUnbook)
				} else {
					if resUnbook.Successful == true {
						fmt.Printf("Deleted reservation '%v' \n", reservation.Id)
					}
				}
			}
		}
		fmt.Printf("Task '%v' finished @ %v \n", "checkForExpiredReservations", time.Now())
	}
}
