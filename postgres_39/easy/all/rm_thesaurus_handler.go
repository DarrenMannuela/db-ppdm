package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmThesaurus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_thesaurus")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_thesaurus

	for rows.Next() {
		var rm_thesaurus dto.Rm_thesaurus
		if err := rows.Scan(&rm_thesaurus.Thesaurus_id, &rm_thesaurus.Active_ind, &rm_thesaurus.Effective_date, &rm_thesaurus.Expiry_date, &rm_thesaurus.Owner_ba_id, &rm_thesaurus.Ppdm_guid, &rm_thesaurus.Remark, &rm_thesaurus.Source, &rm_thesaurus.Source_document_id, &rm_thesaurus.Thesaurus_name, &rm_thesaurus.Thesaurus_version, &rm_thesaurus.Row_changed_by, &rm_thesaurus.Row_changed_date, &rm_thesaurus.Row_created_by, &rm_thesaurus.Row_created_date, &rm_thesaurus.Row_effective_date, &rm_thesaurus.Row_expiry_date, &rm_thesaurus.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_thesaurus to result
		result = append(result, rm_thesaurus)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmThesaurus(c *fiber.Ctx) error {
	var rm_thesaurus dto.Rm_thesaurus

	setDefaults(&rm_thesaurus)

	if err := json.Unmarshal(c.Body(), &rm_thesaurus); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_thesaurus values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	rm_thesaurus.Row_created_date = formatDate(rm_thesaurus.Row_created_date)
	rm_thesaurus.Row_changed_date = formatDate(rm_thesaurus.Row_changed_date)
	rm_thesaurus.Effective_date = formatDateString(rm_thesaurus.Effective_date)
	rm_thesaurus.Expiry_date = formatDateString(rm_thesaurus.Expiry_date)
	rm_thesaurus.Row_effective_date = formatDateString(rm_thesaurus.Row_effective_date)
	rm_thesaurus.Row_expiry_date = formatDateString(rm_thesaurus.Row_expiry_date)






	rows, err := stmt.Exec(rm_thesaurus.Thesaurus_id, rm_thesaurus.Active_ind, rm_thesaurus.Effective_date, rm_thesaurus.Expiry_date, rm_thesaurus.Owner_ba_id, rm_thesaurus.Ppdm_guid, rm_thesaurus.Remark, rm_thesaurus.Source, rm_thesaurus.Source_document_id, rm_thesaurus.Thesaurus_name, rm_thesaurus.Thesaurus_version, rm_thesaurus.Row_changed_by, rm_thesaurus.Row_changed_date, rm_thesaurus.Row_created_by, rm_thesaurus.Row_created_date, rm_thesaurus.Row_effective_date, rm_thesaurus.Row_expiry_date, rm_thesaurus.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmThesaurus(c *fiber.Ctx) error {
	var rm_thesaurus dto.Rm_thesaurus

	setDefaults(&rm_thesaurus)

	if err := json.Unmarshal(c.Body(), &rm_thesaurus); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_thesaurus.Thesaurus_id = id

    if rm_thesaurus.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_thesaurus set active_ind = :1, effective_date = :2, expiry_date = :3, owner_ba_id = :4, ppdm_guid = :5, remark = :6, source = :7, source_document_id = :8, thesaurus_name = :9, thesaurus_version = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where thesaurus_id = :18")
	if err != nil {
		return err
	}

	rm_thesaurus.Row_changed_date = formatDate(rm_thesaurus.Row_changed_date)
	rm_thesaurus.Effective_date = formatDateString(rm_thesaurus.Effective_date)
	rm_thesaurus.Expiry_date = formatDateString(rm_thesaurus.Expiry_date)
	rm_thesaurus.Row_effective_date = formatDateString(rm_thesaurus.Row_effective_date)
	rm_thesaurus.Row_expiry_date = formatDateString(rm_thesaurus.Row_expiry_date)






	rows, err := stmt.Exec(rm_thesaurus.Active_ind, rm_thesaurus.Effective_date, rm_thesaurus.Expiry_date, rm_thesaurus.Owner_ba_id, rm_thesaurus.Ppdm_guid, rm_thesaurus.Remark, rm_thesaurus.Source, rm_thesaurus.Source_document_id, rm_thesaurus.Thesaurus_name, rm_thesaurus.Thesaurus_version, rm_thesaurus.Row_changed_by, rm_thesaurus.Row_changed_date, rm_thesaurus.Row_created_by, rm_thesaurus.Row_effective_date, rm_thesaurus.Row_expiry_date, rm_thesaurus.Row_quality, rm_thesaurus.Thesaurus_id)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchRmThesaurus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_thesaurus set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where thesaurus_id = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteRmThesaurus(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_thesaurus dto.Rm_thesaurus
	rm_thesaurus.Thesaurus_id = id

	stmt, err := db.Prepare("delete from rm_thesaurus where thesaurus_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_thesaurus.Thesaurus_id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


