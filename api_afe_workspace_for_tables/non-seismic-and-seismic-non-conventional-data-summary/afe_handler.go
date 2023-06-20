package apiv1

import (
	"encoding/json"
	"fmt"

	dto "github.com/DarrenMannuela/svc-datatype-nonseismicandseismicnonconventionaldatasummary/dto"
	"github.com/gofiber/fiber/v2"
)

func GetAfe(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT * FROM non_seismic_and_seismic_non_conventional_data_summary_table_afe")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Afe

	for rows.Next() {
		var afe dto.Afe
		if err := rows.Scan(&afe.Afe_number, &afe.Workspace_name, &afe.Kkks_name, &afe.Working_area, &afe.Submission_type, &afe.Data_type); err != nil {
			return err
		}

		// Append r_source to result
		result = append(result, afe)
	}
	// Return result in JSON format
	return c.JSON(result)

}

func GetAfeByNum(c *fiber.Ctx) error {
	afeNumber := c.Params("afe")
	fmt.Println(afeNumber)
	rows, err := db.Query("SELECT * FROM non_seismic_and_seismic_non_conventional_data_summary_table_afe WHERE afe_number = :1", afeNumber)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Afe

	for rows.Next() {
		var afe dto.Afe
		if err := rows.Scan(&afe.Afe_number, &afe.Workspace_name, &afe.Kkks_name, &afe.Working_area, &afe.Submission_type, &afe.Data_type); err != nil {
			return err
		}

		// Append r_source to result
		result = append(result, afe)
	}
	// Return result in JSON format
	return c.JSON(result)

}

func SetAfe(c *fiber.Ctx) error {
	var afe dto.Afe
	setDefaults(&afe)
	afe.Afe_number = generateRandomInt(1, 100)

	if err := json.Unmarshal(c.Body(), &afe); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	afe.Data_type = "Non Seismic And Seismic Non Conventional Data Summary"

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	var submissionTypeExist string
	_ = tx.QueryRow(`SELECT submission_type FROM submission WHERE submission_type = :1`, afe.Submission_type).Scan(&submissionTypeExist)

	if submissionTypeExist == "" {
		_, err = tx.Exec(`INSERT INTO submission (submission_type) VALUES (:1)`, afe.Submission_type)
		if err != nil {
			tx.Rollback()
			fmt.Println("SUBMISSION")
			return err

		}
	}
	_, err = tx.Exec(`INSERT INTO non_seismic_and_seismic_non_conventional_data_summary_table_afe (afe_number, workspace_name, kkks_name, working_area, submission_type, data_type) VALUES (:1, :2, :3, :4, :5, :6)`, afe.Afe_number, afe.Workspace_name, afe.Kkks_name, afe.Working_area, afe.Submission_type, afe.Data_type)
	if err != nil {
		tx.Rollback()
		fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE_AFE")
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)

}

func PutAfe(c *fiber.Ctx) error {
	afeNumber := c.Params("afe")

	var afe dto.Afe
	setDefaults(&afe)
	if err := json.Unmarshal(c.Body(), &afe); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	afe.Data_type = "Non Seismic And Seismic Non Conventional Data Summary"

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	var submissionTypeExist string
	var submissionDoesntExist bool
	err = tx.QueryRow(`SELECT submission_type FROM submission WHERE submission_type = :1`, afe.Submission_type).Scan(&submissionTypeExist)

	if err != nil {
		submissionDoesntExist = true
	}
	if submissionDoesntExist {
		_, err = tx.Exec(`INSERT INTO submission (submission_type) VALUES (:1)`, afe.Submission_type)
		if err != nil {
			fmt.Println("SUBMISSION")
			return err
		}
	}
	_, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_data_summary_table_afe SET workspace_name = :1, kkks_name = :2, working_area = :3, submission_type = :4, data_type = :5 WHERE afe_number = :6`,
		afe.Workspace_name, afe.Kkks_name, afe.Working_area, afe.Submission_type, afe.Data_type, afeNumber)
	if err != nil {
		tx.Rollback()
		fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE_AFE")
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)

}

func DeleteAfe(c *fiber.Ctx) error {
	afeNumber := c.Params("afe")

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE_AFE")
		return c.Status(500).SendString(err.Error())
	}

	rows, err := tx.Query(`SELECT non_seismic_and_seismic_non_conventional_data_summary_id FROM non_seismic_and_seismic_non_conventional_data_summary_table_workspace WHERE afe_number = :1`, afeNumber)
	if err != nil {
		tx.Rollback()
		fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE_WORKSPACE")
		return err
	}

	var non_seismic_and_seismic_non_conventional_data_summaryIds []string

	for rows.Next() {
		var non_seismic_and_seismic_non_conventional_data_summaryId string
		if err := rows.Scan(&non_seismic_and_seismic_non_conventional_data_summaryId); err != nil {
			return err
		}

		non_seismic_and_seismic_non_conventional_data_summaryIds = append(non_seismic_and_seismic_non_conventional_data_summaryIds, non_seismic_and_seismic_non_conventional_data_summaryId)

	}

	_, err = tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_data_summary_table_workspace WHERE afe_number = :1`, afeNumber)
	if err != nil {
		tx.Rollback()
		fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE_WORKSPACE")
		return err
	}

	for i := 0; i < len(non_seismic_and_seismic_non_conventional_data_summaryIds); i++ {
		_, err = tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_data_summary_table WHERE id = :1`, non_seismic_and_seismic_non_conventional_data_summaryIds[i])
		if err != nil {
			tx.Rollback()
			fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE")
			return err
		}
	}

	delRows, err := tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_data_summary_table_afe WHERE afe_number = :1`, afeNumber)
	if err != nil {
		tx.Rollback()
		fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE_AFE")
		return err
	}

	affectedRows, err := delRows.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows > 0 {
		fmt.Printf("%d row(s) deleted\n", affectedRows)
	} else {
		fmt.Println("No rows deleted")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)

}

func PatchWorkspaceAfe(c *fiber.Ctx) error {
	afeNumber := c.Params("afe")

	var afe dto.Afe
	if err := json.Unmarshal(c.Body(), &afe); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}

	if afe.Workspace_name != nil {

		_, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_data_summary_table_afe SET workspace_name = :1 WHERE afe_number = :2`, afe.Workspace_name, afeNumber)

		if err != nil {
			tx.Rollback()
			fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE_AFE")
			return err
		}
	}
	if afe.Kkks_name != nil {

		_, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_data_summary_table_afe SET kkks_name = :1 WHERE afe_number = :2`, afe.Kkks_name, afeNumber)

		if err != nil {
			tx.Rollback()
			fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE_AFE")
			return err
		}
	}
	if afe.Working_area != nil {

		_, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_data_summary_table_afe SET working_area = :1 WHERE afe_number = :2`, afe.Working_area, afeNumber)

		if err != nil {
			tx.Rollback()
			fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE_AFE")
			return err
		}
	}
	if afe.Submission_type != nil {
		var submissionTypeExist string
		var submissionDoesntExist bool
		err = tx.QueryRow(`SELECT submission_type FROM submission WHERE submission_type = :1`, afe.Submission_type).Scan(&submissionTypeExist)

		if err != nil {
			submissionDoesntExist = true
		}
		if submissionDoesntExist {
			_, err = tx.Exec(`INSERT INTO submission (submission_type) VALUES (:1)`, afe.Submission_type)
			if err != nil {
				fmt.Println("SUBMISSION")
				return err
			}
		}

		_, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_data_summary_table_afe SET submission_type = :1 WHERE afe_number = :2`, afe.Submission_type, afeNumber)

		if err != nil {
			tx.Rollback()
			fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE_AFE")
			return err
		}
	}
	if afe.Data_type != "" {

		_, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_data_summary_table_afe SET data_type = :1 WHERE afe_number = :2`, afe.Data_type, afeNumber)

		if err != nil {
			tx.Rollback()
			fmt.Println("NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE_AFE")
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
