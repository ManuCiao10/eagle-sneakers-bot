package ws_quicktasking

import (
	"fmt"
	"strconv"

	"github.com/valyala/fastjson"
)

func handleQuicktaskMessage(message []byte) {
	siteId := fastjson.GetString(message, "siteId")
	msku := fastjson.GetString(message, "product_id")
	size := fastjson.GetString(message, "size")

	// proxyGroupId, _ := proxy.GetProxyGroupIDByName("main")

	siteIdInt, err := strconv.Atoi(siteId)
	if err != nil {
		fmt.Println("Failed to convert siteId to int.")
		return
	}
	fmt.Println(message, siteId, msku, size, siteIdInt)

	// switch siteId {
	// case "1":
	// 	for _, accountGroup := range loading.Data.Quicktask.Quicktask[siteIdInt] {
	// 		profileId, err := profile.GetProfileIDByName(accountGroup.ProfileName)
	// 		if err != nil {
	// 			fmt.Println("Quicktask failed: profile not found.")
	// 			return
	// 		}

	// 		taskUUID := task.CreateTask(
	// 			"fuelmonitor",
	// 			msku,
	// 			size,
	// 			profileId,
	// 			proxyGroupId,
	// 			accountGroup.AccountEmail,
	// 			"monitor",
	// 			"fast",
	// 			2000,
	// 		)
	// 		tObject, _ := task.GetTask(taskUUID)
	// 		go taskmngr.RunTask(tObject)
	// 	}
	// case "0":
	// 	for _, accountGroup := range loading.Data.QuicktaskGroups[siteIdInt] {
	// 		profileId, err := profile.GetProfileIDByName(accountGroup.ProfileName)
	// 		if err != nil {
	// 			fmt.Println("Quicktask failed: profile not found.")
	// 			return
	// 		}

	// 		monitorTaskUUID := task.CreateTask(
	// 			"athletesfootmonitor",
	// 			msku,
	// 			size,
	// 			profileId,
	// 			proxyGroupId,
	// 			accountGroup.AccountEmail,
	// 			"monitor",
	// 			"msku",
	// 			2000,
	// 		)

	// 		checkoutTaskUUID := task.CreateTask(
	// 			"athletesfoot",
	// 			msku,
	// 			size,
	// 			profileId,
	// 			proxyGroupId,
	// 			accountGroup.AccountEmail,
	// 			"checkout",
	// 			"msku",
	// 			2000,
	// 		)

	// 		monitorObject, _ := task.GetTask(monitorTaskUUID)
	// 		checkoutObject, _ := task.GetTask(checkoutTaskUUID)
	// 		go taskmngr.RunTask(monitorObject)
	// 		go taskmngr.RunTask(checkoutObject)
	// 	}
	// case "3":
	// 	for _, accountGroup := range loading.Data.QuicktaskGroups[siteIdInt] {
	// 		profileId, err := profile.GetProfileIDByName(accountGroup.ProfileName)
	// 		if err != nil {
	// 			fmt.Println("Quicktask failed: profile not found.")
	// 			return
	// 		}

	// 		monitorTaskUUID := task.CreateTask(
	// 			"buzzsneakersmonitor",
	// 			msku,
	// 			size,
	// 			profileId,
	// 			proxyGroupId,
	// 			accountGroup.AccountEmail,
	// 			"monitor",
	// 			"msku",
	// 			2000,
	// 		)

	// 		checkoutTaskUUID := task.CreateTask(
	// 			"buzzsneakers",
	// 			msku,
	// 			size,
	// 			profileId,
	// 			proxyGroupId,
	// 			accountGroup.AccountEmail,
	// 			"checkout",
	// 			"normal",
	// 			2000,
	// 		)

	// 		monitorObject, _ := task.GetTask(monitorTaskUUID)
	// 		checkoutObject, _ := task.GetTask(checkoutTaskUUID)
	// 		go taskmngr.RunTask(checkoutObject)
	// 		time.Sleep(500 * time.Millisecond)
	// 		go taskmngr.RunTask(monitorObject)
	// 	}
	// }
}
