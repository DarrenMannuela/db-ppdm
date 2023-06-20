package apiv1

import (
	"encoding/json"
	"fmt"

	dto "github.com/DarrenMannuela/svc-datatype-projectfilestoredinmedia/dto"
	"github.com/gofiber/fiber/v2"
)

func GetAfe(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT * FROM project_file_stored_in_media_table_afe")
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
	rows, err := db.Query("SELECT * FROM project_file_stored_in_media_table_afe WHERE afe_number = :1", afeNumber)
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

	afe.Data_type = "Project File Stored In Media"

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
	_, err = tx.Exec(`INSERT INTO project_file_stored_in_media_table_afe (afe_number, workspace_name, kkks_name, working_area, submission_type, data_type) VALUES (:1, :2, :3, :4, :5, :6)`, afe.Afe_number, afe.Workspace_name, afe.Kkks_name, afe.Working_area, afe.Submission_type, afe.Data_type)
	if err != nil {
		tx.Rollback()
		fmt.Println("PROJECT_FILE_STORED_IN_MEDIA_TABLE_AFE")
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

	afe.Data_type = "Project File Stored In Media"

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
	_, err = tx.Exec(`UPDATE project_file_stored_in_media_table_afe SET workspace_name = :1, kkks_name = :2, working_area = :3, submission_type = :4, data_type = :5 WHERE afe_number = :6`,
		afe.Workspace_name, afe.Kkks_name, afe.Working_area, afe.Submission_type, afe.Data_type, afeNumber)
	if err != nil {
		tx.Rollback()
		fmt.Println("PROJECT_FILE_STORED_IN_MEDIA_TABLE_AFE")
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
		fmt.Println("PROJECT_FILE_STORED_IN_MEDIA_TABLE_AFE")
		return c.Status(500).SendString(err.Error())
	}

	rows, err := tx.Query(`SELECT project_file_stored_in_media_id FROM project_file_stored_in_media_table_workspace WHERE afe_number = :1`, afeNumber)
	if err != nil {
		tx.Rollback()
		fmt.Println("PROJECT_FILE_STORED_IN_MEDIA_TABLE_WORKSPACE")
		return err
	}

	var project_file_stored_in_mediaIds []string

	for rows.Next() {
		var project_file_stored_in_mediaId string
		if err := rows.Scan(&project_file_stored_in_mediaId); err != nil {
			return err
		}

		project_file_stored_in_mediaIds = append(project_file_stored_in_mediaIds, project_file_stored_in_mediaId)

	}

	_, err = tx.Exec(`DELETE FROM project_file_stored_in_media_table_workspace WHERE afe_number = :1`, afeNumber)
	if err != nil {
		tx.Rollback()
		fmt.Println("PROJECT_FILE_STORED_IN_MEDIA_TABLE_WORKSPACE")
		return err
	}

	for i := 0; i < len(project_file_stored_in_mediaIds); i++ {
		_, err = tx.Exec(`DELETE FROM project_file_stored_in_media_table WHERE id = :1`, project_file_stored_in_mediaIds[i])
		if err != nil {
			tx.Rollback()
			fmt.Println("PROJECT_FILE_STORED_IN_MEDIA_TABLE")
			return err
		}
	}

	delRows, err := tx.Exec(`DELETE FROM project_file_stored_in_media_table_afe WHERE afe_number = :1`, afeNumber)
	if err != nil {
		tx.Rollback()
		fmt.Println("PROJECT_FILE_STORED_IN_MEDIA_TABLE_AFE")
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

		_, err = tx.Exec(`UPDATE project_file_stored_in_media_table_afe SET workspace_name = :1 WHERE afe_number = :2`, afe.Workspace_name, afeNumber)

		if err != nil {
			tx.Rollback()
			fmt.Println("PROJECT_FILE_STORED_IN_MEDIA_TABLE_AFE")
			return err
		}
	}
	if afe.Kkks_name != nil {

		_, err = tx.Exec(`UPDATE project_file_stored_in_media_table_afe SET kkks_name = :1 WHERE afe_number = :2`, afe.Kkks_name, afeNumber)

		if err != nil {
			tx.Rollback()
			fmt.Println("PROJECT_FILE_STORED_IN_MEDIA_TABLE_AFE")
			return err
		}
	}
	if afe.Working_area != nil {

		_, err = tx.Exec(`UPDATE project_file_stored_in_media_table_afe SET working_area = :1 WHERE afe_number = :2`, afe.Working_area, afeNumber)

		if err != nil {
			tx.Rollback()
			fmt.Println("PROJECT_FILE_STORED_IN_MEDIA_TABLE_AFE")
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

		_, err = tx.Exec(`UPDATE project_file_stored_in_media_table_afe SET submission_type = :1 WHERE afe_number = :2`, afe.Submission_type, afeNumber)

		if err != nil {
			tx.Rollback()
			fmt.Println("PROJECT_FILE_STORED_IN_MEDIA_TABLE_AFE")
			return err
		}
	}
	if afe.Data_type != "" {

		_, err = tx.Exec(`UPDATE project_file_stored_in_media_table_afe SET data_type = :1 WHERE afe_number = :2`, afe.Data_type, afeNumber)

		if err != nil {
			tx.Rollback()
			fmt.Println("PROJECT_FILE_STORED_IN_MEDIA_TABLE_AFE")
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
