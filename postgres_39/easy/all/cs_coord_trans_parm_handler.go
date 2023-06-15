package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCsCoordTransParm(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cs_coord_trans_parm")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cs_coord_trans_parm

	for rows.Next() {
		var cs_coord_trans_parm dto.Cs_coord_trans_parm
		if err := rows.Scan(&cs_coord_trans_parm.Transform_type, &cs_coord_trans_parm.Parameter_id, &cs_coord_trans_parm.Active_ind, &cs_coord_trans_parm.Digital_output, &cs_coord_trans_parm.Effective_date, &cs_coord_trans_parm.Expiry_date, &cs_coord_trans_parm.Ppdm_guid, &cs_coord_trans_parm.Remark, &cs_coord_trans_parm.Source, &cs_coord_trans_parm.Source_document_id, &cs_coord_trans_parm.Transform_parm, &cs_coord_trans_parm.Row_changed_by, &cs_coord_trans_parm.Row_changed_date, &cs_coord_trans_parm.Row_created_by, &cs_coord_trans_parm.Row_created_date, &cs_coord_trans_parm.Row_effective_date, &cs_coord_trans_parm.Row_expiry_date, &cs_coord_trans_parm.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cs_coord_trans_parm to result
		result = append(result, cs_coord_trans_parm)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCsCoordTransParm(c *fiber.Ctx) error {
	var cs_coord_trans_parm dto.Cs_coord_trans_parm

	setDefaults(&cs_coord_trans_parm)

	if err := json.Unmarshal(c.Body(), &cs_coord_trans_parm); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cs_coord_trans_parm values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	cs_coord_trans_parm.Row_created_date = formatDate(cs_coord_trans_parm.Row_created_date)
	cs_coord_trans_parm.Row_changed_date = formatDate(cs_coord_trans_parm.Row_changed_date)
	cs_coord_trans_parm.Effective_date = formatDateString(cs_coord_trans_parm.Effective_date)
	cs_coord_trans_parm.Expiry_date = formatDateString(cs_coord_trans_parm.Expiry_date)
	cs_coord_trans_parm.Row_effective_date = formatDateString(cs_coord_trans_parm.Row_effective_date)
	cs_coord_trans_parm.Row_expiry_date = formatDateString(cs_coord_trans_parm.Row_expiry_date)






	rows, err := stmt.Exec(cs_coord_trans_parm.Transform_type, cs_coord_trans_parm.Parameter_id, cs_coord_trans_parm.Active_ind, cs_coord_trans_parm.Digital_output, cs_coord_trans_parm.Effective_date, cs_coord_trans_parm.Expiry_date, cs_coord_trans_parm.Ppdm_guid, cs_coord_trans_parm.Remark, cs_coord_trans_parm.Source, cs_coord_trans_parm.Source_document_id, cs_coord_trans_parm.Transform_parm, cs_coord_trans_parm.Row_changed_by, cs_coord_trans_parm.Row_changed_date, cs_coord_trans_parm.Row_created_by, cs_coord_trans_parm.Row_created_date, cs_coord_trans_parm.Row_effective_date, cs_coord_trans_parm.Row_expiry_date, cs_coord_trans_parm.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCsCoordTransParm(c *fiber.Ctx) error {
	var cs_coord_trans_parm dto.Cs_coord_trans_parm

	setDefaults(&cs_coord_trans_parm)

	if err := json.Unmarshal(c.Body(), &cs_coord_trans_parm); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cs_coord_trans_parm.Transform_type = id

    if cs_coord_trans_parm.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cs_coord_trans_parm set parameter_id = :1, active_ind = :2, digital_output = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, source_document_id = :9, transform_parm = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where transform_type = :18")
	if err != nil {
		return err
	}

	cs_coord_trans_parm.Row_changed_date = formatDate(cs_coord_trans_parm.Row_changed_date)
	cs_coord_trans_parm.Effective_date = formatDateString(cs_coord_trans_parm.Effective_date)
	cs_coord_trans_parm.Expiry_date = formatDateString(cs_coord_trans_parm.Expiry_date)
	cs_coord_trans_parm.Row_effective_date = formatDateString(cs_coord_trans_parm.Row_effective_date)
	cs_coord_trans_parm.Row_expiry_date = formatDateString(cs_coord_trans_parm.Row_expiry_date)






	rows, err := stmt.Exec(cs_coord_trans_parm.Parameter_id, cs_coord_trans_parm.Active_ind, cs_coord_trans_parm.Digital_output, cs_coord_trans_parm.Effective_date, cs_coord_trans_parm.Expiry_date, cs_coord_trans_parm.Ppdm_guid, cs_coord_trans_parm.Remark, cs_coord_trans_parm.Source, cs_coord_trans_parm.Source_document_id, cs_coord_trans_parm.Transform_parm, cs_coord_trans_parm.Row_changed_by, cs_coord_trans_parm.Row_changed_date, cs_coord_trans_parm.Row_created_by, cs_coord_trans_parm.Row_effective_date, cs_coord_trans_parm.Row_expiry_date, cs_coord_trans_parm.Row_quality, cs_coord_trans_parm.Transform_type)
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

func PatchCsCoordTransParm(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cs_coord_trans_parm set "
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
	query += " where transform_type = :id"

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

func DeleteCsCoordTransParm(c *fiber.Ctx) error {
	id := c.Params("id")
	var cs_coord_trans_parm dto.Cs_coord_trans_parm
	cs_coord_trans_parm.Transform_type = id

	stmt, err := db.Prepare("delete from cs_coord_trans_parm where transform_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cs_coord_trans_parm.Transform_type)
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


