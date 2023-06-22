package apiv1

import (
	"encoding/json"
	"fmt"
	"log"

	dto "github.com/DarrenMannuela/svc-datatype-nonseismicandseismicnonconventionalreport/dto"

	"github.com/gofiber/fiber/v2"
)

func GetWorkspace(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT * FROM non_seismic_and_seismic_non_conventional_report_table_workspace")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Workspace

	for rows.Next() {
		var workspace dto.Workspace
		if err := rows.Scan(&workspace.Id, &workspace.Afe_number, &workspace.Non_seismic_and_seismic_non_conventional_report_id); err != nil {
			return err // Exit if we get an error
		}

		// Append r_source to result
		result = append(result, workspace)
	}
	// Return result in JSON format
	return c.JSON(result)

}

func GetWorkspaceByAfe(c *fiber.Ctx) error {
	id := c.Params("afe")

	var data []dto.Workspace

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	rows, err := tx.Query("SELECT * FROM non_seismic_and_seismic_non_conventional_report_table_workspace WHERE afe_number = :1", id)
	if err != nil {
		return err
	}

	for rows.Next() {
		var curRow dto.Workspace
		if err := rows.Scan(&curRow.Id, &curRow.Afe_number, &curRow.Non_seismic_and_seismic_non_conventional_report_id); err != nil {
			return err // Exit if we get an error
		}

		data = append(data, curRow)

	}
	return c.JSON(data)

}

func SetWorkspace(c *fiber.Ctx) error {
	id := generateRandomInt(0, 1000000)
	var workspace dto.Workspace
	setDefaults(&workspace)
	if err := json.Unmarshal(c.Body(), &workspace); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	workspace.Id = id

	var submission dto.Submission
	setDefaults(&submission)

	if err := json.Unmarshal(c.Body(), &submission); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var afe dto.Afe
	setDefaults(&afe)

	if err := json.Unmarshal(c.Body(), &afe); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			err := fmt.Errorf("SetWorkspace panicked: %v", p)
			log.Println(err)
			panic(err)
		} else if err != nil {
			tx.Rollback()
			log.Println(err)
		} else {
			tx.Commit()
		}
	}()

	var Non_seismic_and_seismic_non_conventional_report_id_exist string
	err = tx.QueryRow("SELECT id FROM non_seismic_and_seismic_non_conventional_report_table where id = :1", workspace.Non_seismic_and_seismic_non_conventional_report_id).Scan(&Non_seismic_and_seismic_non_conventional_report_id_exist)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO non_seismic_and_seismic_non_conventional_report_table_workspace (afe_number, non_seismic_and_seismic_non_conventional_report_id) VALUES (:2, :3)`, workspace.Afe_number, workspace.Non_seismic_and_seismic_non_conventional_report_id)
	if err != nil {
		tx.Rollback()
		fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_TABLE_WORKSPACE")
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func PutWorkspace(c *fiber.Ctx) error {
	id := c.Params("id")

	var workspace dto.Workspace
	setDefaults(&workspace)
	if err := json.Unmarshal(c.Body(), &workspace); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	_, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table_workspace SET afe_number = :1, non_seismic_and_seismic_non_conventional_report_id = :2 WHERE id = :3`,
		workspace.Afe_number, workspace.Non_seismic_and_seismic_non_conventional_report_id, id)
	if err != nil {
		tx.Rollback()
		fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_TABLE_WORKSPACE")
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)

}

func DeleteWorkspace(c *fiber.Ctx) error {
	id := c.Params("id")

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	_, err = tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_report_table_workspace WHERE non_seismic_and_seismic_non_conventional_report_id = :1`, id)
	if err != nil {
		fmt.Println("DEL NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_TABLE_WORKSPACE")
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	_, err = tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_report_table WHERE id = :1`, id)
	if err != nil {
		fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_TABLE")
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func PatchWorkspace(c *fiber.Ctx) error {
	id := c.Params("id")

	var workspace dto.Workspace

	if err := json.Unmarshal(c.Body(), &workspace); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	if workspace.Afe_number != nil {

		_, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table_workspace SET afe_number = :1 WHERE id = :2`,
			workspace.Afe_number, id)
		if err != nil {
			tx.Rollback()
			fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_TABLE_WORKSPACE")
			return err
		}

	}

	if workspace.Non_seismic_and_seismic_non_conventional_report_id != 0 {

		_, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table_workspace SET non_seismic_and_seismic_non_conventional_report_id = :1 WHERE id = :2`,
			workspace.Non_seismic_and_seismic_non_conventional_report_id, id)
		if err != nil {
			tx.Rollback()
			fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_WORKSPACE")
			return err
		}

	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)

}
