package apiv1

import (
	"encoding/json"
	"fmt"
	"log"

	dto "github.com/DarrenMannuela/svc-datatype-digitalwelllog/dto"

	"github.com/gofiber/fiber/v2"
)

func GetWorkspace(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT * FROM digital_well_log_table_workspace")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Workspace

	for rows.Next() {
		var workspace dto.Workspace
		if err := rows.Scan(&workspace.Id, &workspace.Afe_number, &workspace.Print_well_report_id); err != nil {
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

	rows, err := tx.Query("SELECT * FROM digital_well_log_table_workspace WHERE afe_number = :1", id)
	if err != nil {
		return err
	}

	for rows.Next() {
		var curRow dto.Workspace
		if err := rows.Scan(&curRow.Id, &curRow.Afe_number, &curRow.Print_well_report_id); err != nil {
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

	var Print_well_report_id_exist string
	err = tx.QueryRow("SELECT id FROM digital_well_log_table where id = :1", workspace.Print_well_report_id).Scan(&Print_well_report_id_exist)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO digital_well_log_table_workspace (afe_number, Print_well_report_id) VALUES (:2, :3)`, workspace.Afe_number, workspace.Print_well_report_id)
	if err != nil {
		tx.Rollback()
		fmt.Println("DIGITAL_WELL_LOG_TABLE_WORKSPACE")
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

	_, err = tx.Exec(`UPDATE digital_well_log_table_workspace SET afe_number = :1, Print_well_report_id = :2 WHERE id = :3`,
		workspace.Afe_number, workspace.Print_well_report_id, id)
	if err != nil {
		tx.Rollback()
		fmt.Println("DIGITAL_WELL_LOG_TABLE_WORKSPACE")
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

	_, err = tx.Exec(`DELETE FROM digital_well_log_table_workspace WHERE Print_well_report_id = :1`, id)
	if err != nil {
		fmt.Println("DEL DIGITAL_WELL_LOG_TABLE_WORKSPACE")
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	_, err = tx.Exec(`DELETE FROM digital_well_log_table WHERE id = :1`, id)
	if err != nil {
		fmt.Println("DIGITAL_WELL_LOG_TABLE")
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

		_, err = tx.Exec(`UPDATE digital_well_log_table_workspace SET afe_number = :1 WHERE id = :2`,
			workspace.Afe_number, id)
		if err != nil {
			tx.Rollback()
			fmt.Println("DIGITAL_WELL_LOG_TABLE_WORKSPACE")
			return err
		}

	}

	if workspace.Print_well_report_id != 0 {

		_, err = tx.Exec(`UPDATE digital_well_log_table_workspace SET Print_well_report_id = :1 WHERE id = :2`,
			workspace.Print_well_report_id, id)
		if err != nil {
			tx.Rollback()
			fmt.Println("DIGITAL_WELL_LOG_WORKSPACE")
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
