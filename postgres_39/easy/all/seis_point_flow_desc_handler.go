package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisPointFlowDesc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_point_flow_desc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_point_flow_desc

	for rows.Next() {
		var seis_point_flow_desc dto.Seis_point_flow_desc
		if err := rows.Scan(&seis_point_flow_desc.Seis_set_subtype, &seis_point_flow_desc.Seis_set_id, &seis_point_flow_desc.Seis_point_id, &seis_point_flow_desc.Flow_id, &seis_point_flow_desc.Active_ind, &seis_point_flow_desc.Description, &seis_point_flow_desc.Description_type, &seis_point_flow_desc.Effective_date, &seis_point_flow_desc.Expiry_date, &seis_point_flow_desc.Ppdm_guid, &seis_point_flow_desc.Remark, &seis_point_flow_desc.Source, &seis_point_flow_desc.Row_changed_by, &seis_point_flow_desc.Row_changed_date, &seis_point_flow_desc.Row_created_by, &seis_point_flow_desc.Row_created_date, &seis_point_flow_desc.Row_effective_date, &seis_point_flow_desc.Row_expiry_date, &seis_point_flow_desc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_point_flow_desc to result
		result = append(result, seis_point_flow_desc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisPointFlowDesc(c *fiber.Ctx) error {
	var seis_point_flow_desc dto.Seis_point_flow_desc

	setDefaults(&seis_point_flow_desc)

	if err := json.Unmarshal(c.Body(), &seis_point_flow_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_point_flow_desc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	seis_point_flow_desc.Row_created_date = formatDate(seis_point_flow_desc.Row_created_date)
	seis_point_flow_desc.Row_changed_date = formatDate(seis_point_flow_desc.Row_changed_date)
	seis_point_flow_desc.Effective_date = formatDateString(seis_point_flow_desc.Effective_date)
	seis_point_flow_desc.Expiry_date = formatDateString(seis_point_flow_desc.Expiry_date)
	seis_point_flow_desc.Row_effective_date = formatDateString(seis_point_flow_desc.Row_effective_date)
	seis_point_flow_desc.Row_expiry_date = formatDateString(seis_point_flow_desc.Row_expiry_date)






	rows, err := stmt.Exec(seis_point_flow_desc.Seis_set_subtype, seis_point_flow_desc.Seis_set_id, seis_point_flow_desc.Seis_point_id, seis_point_flow_desc.Flow_id, seis_point_flow_desc.Active_ind, seis_point_flow_desc.Description, seis_point_flow_desc.Description_type, seis_point_flow_desc.Effective_date, seis_point_flow_desc.Expiry_date, seis_point_flow_desc.Ppdm_guid, seis_point_flow_desc.Remark, seis_point_flow_desc.Source, seis_point_flow_desc.Row_changed_by, seis_point_flow_desc.Row_changed_date, seis_point_flow_desc.Row_created_by, seis_point_flow_desc.Row_created_date, seis_point_flow_desc.Row_effective_date, seis_point_flow_desc.Row_expiry_date, seis_point_flow_desc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisPointFlowDesc(c *fiber.Ctx) error {
	var seis_point_flow_desc dto.Seis_point_flow_desc

	setDefaults(&seis_point_flow_desc)

	if err := json.Unmarshal(c.Body(), &seis_point_flow_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_point_flow_desc.Seis_set_subtype = id

    if seis_point_flow_desc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_point_flow_desc set seis_set_id = :1, seis_point_id = :2, flow_id = :3, active_ind = :4, description = :5, description_type = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where seis_set_subtype = :19")
	if err != nil {
		return err
	}

	seis_point_flow_desc.Row_changed_date = formatDate(seis_point_flow_desc.Row_changed_date)
	seis_point_flow_desc.Effective_date = formatDateString(seis_point_flow_desc.Effective_date)
	seis_point_flow_desc.Expiry_date = formatDateString(seis_point_flow_desc.Expiry_date)
	seis_point_flow_desc.Row_effective_date = formatDateString(seis_point_flow_desc.Row_effective_date)
	seis_point_flow_desc.Row_expiry_date = formatDateString(seis_point_flow_desc.Row_expiry_date)






	rows, err := stmt.Exec(seis_point_flow_desc.Seis_set_id, seis_point_flow_desc.Seis_point_id, seis_point_flow_desc.Flow_id, seis_point_flow_desc.Active_ind, seis_point_flow_desc.Description, seis_point_flow_desc.Description_type, seis_point_flow_desc.Effective_date, seis_point_flow_desc.Expiry_date, seis_point_flow_desc.Ppdm_guid, seis_point_flow_desc.Remark, seis_point_flow_desc.Source, seis_point_flow_desc.Row_changed_by, seis_point_flow_desc.Row_changed_date, seis_point_flow_desc.Row_created_by, seis_point_flow_desc.Row_effective_date, seis_point_flow_desc.Row_expiry_date, seis_point_flow_desc.Row_quality, seis_point_flow_desc.Seis_set_subtype)
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

func PatchSeisPointFlowDesc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_point_flow_desc set "
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
	query += " where seis_set_subtype = :id"

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

func DeleteSeisPointFlowDesc(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_point_flow_desc dto.Seis_point_flow_desc
	seis_point_flow_desc.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_point_flow_desc where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_point_flow_desc.Seis_set_subtype)
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


