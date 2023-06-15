package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogDictCrvCls(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_dict_crv_cls")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_dict_crv_cls

	for rows.Next() {
		var well_log_dict_crv_cls dto.Well_log_dict_crv_cls
		if err := rows.Scan(&well_log_dict_crv_cls.Curve_class_id, &well_log_dict_crv_cls.Dictionary_id, &well_log_dict_crv_cls.Dict_curve_id, &well_log_dict_crv_cls.Active_ind, &well_log_dict_crv_cls.Effective_date, &well_log_dict_crv_cls.Expiry_date, &well_log_dict_crv_cls.Ppdm_guid, &well_log_dict_crv_cls.Remark, &well_log_dict_crv_cls.Source, &well_log_dict_crv_cls.Row_changed_by, &well_log_dict_crv_cls.Row_changed_date, &well_log_dict_crv_cls.Row_created_by, &well_log_dict_crv_cls.Row_created_date, &well_log_dict_crv_cls.Row_effective_date, &well_log_dict_crv_cls.Row_expiry_date, &well_log_dict_crv_cls.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_dict_crv_cls to result
		result = append(result, well_log_dict_crv_cls)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogDictCrvCls(c *fiber.Ctx) error {
	var well_log_dict_crv_cls dto.Well_log_dict_crv_cls

	setDefaults(&well_log_dict_crv_cls)

	if err := json.Unmarshal(c.Body(), &well_log_dict_crv_cls); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_dict_crv_cls values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	well_log_dict_crv_cls.Row_created_date = formatDate(well_log_dict_crv_cls.Row_created_date)
	well_log_dict_crv_cls.Row_changed_date = formatDate(well_log_dict_crv_cls.Row_changed_date)
	well_log_dict_crv_cls.Effective_date = formatDateString(well_log_dict_crv_cls.Effective_date)
	well_log_dict_crv_cls.Expiry_date = formatDateString(well_log_dict_crv_cls.Expiry_date)
	well_log_dict_crv_cls.Row_effective_date = formatDateString(well_log_dict_crv_cls.Row_effective_date)
	well_log_dict_crv_cls.Row_expiry_date = formatDateString(well_log_dict_crv_cls.Row_expiry_date)






	rows, err := stmt.Exec(well_log_dict_crv_cls.Curve_class_id, well_log_dict_crv_cls.Dictionary_id, well_log_dict_crv_cls.Dict_curve_id, well_log_dict_crv_cls.Active_ind, well_log_dict_crv_cls.Effective_date, well_log_dict_crv_cls.Expiry_date, well_log_dict_crv_cls.Ppdm_guid, well_log_dict_crv_cls.Remark, well_log_dict_crv_cls.Source, well_log_dict_crv_cls.Row_changed_by, well_log_dict_crv_cls.Row_changed_date, well_log_dict_crv_cls.Row_created_by, well_log_dict_crv_cls.Row_created_date, well_log_dict_crv_cls.Row_effective_date, well_log_dict_crv_cls.Row_expiry_date, well_log_dict_crv_cls.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogDictCrvCls(c *fiber.Ctx) error {
	var well_log_dict_crv_cls dto.Well_log_dict_crv_cls

	setDefaults(&well_log_dict_crv_cls)

	if err := json.Unmarshal(c.Body(), &well_log_dict_crv_cls); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_dict_crv_cls.Curve_class_id = id

    if well_log_dict_crv_cls.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_dict_crv_cls set dictionary_id = :1, dict_curve_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where curve_class_id = :16")
	if err != nil {
		return err
	}

	well_log_dict_crv_cls.Row_changed_date = formatDate(well_log_dict_crv_cls.Row_changed_date)
	well_log_dict_crv_cls.Effective_date = formatDateString(well_log_dict_crv_cls.Effective_date)
	well_log_dict_crv_cls.Expiry_date = formatDateString(well_log_dict_crv_cls.Expiry_date)
	well_log_dict_crv_cls.Row_effective_date = formatDateString(well_log_dict_crv_cls.Row_effective_date)
	well_log_dict_crv_cls.Row_expiry_date = formatDateString(well_log_dict_crv_cls.Row_expiry_date)






	rows, err := stmt.Exec(well_log_dict_crv_cls.Dictionary_id, well_log_dict_crv_cls.Dict_curve_id, well_log_dict_crv_cls.Active_ind, well_log_dict_crv_cls.Effective_date, well_log_dict_crv_cls.Expiry_date, well_log_dict_crv_cls.Ppdm_guid, well_log_dict_crv_cls.Remark, well_log_dict_crv_cls.Source, well_log_dict_crv_cls.Row_changed_by, well_log_dict_crv_cls.Row_changed_date, well_log_dict_crv_cls.Row_created_by, well_log_dict_crv_cls.Row_effective_date, well_log_dict_crv_cls.Row_expiry_date, well_log_dict_crv_cls.Row_quality, well_log_dict_crv_cls.Curve_class_id)
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

func PatchWellLogDictCrvCls(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_dict_crv_cls set "
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
	query += " where curve_class_id = :id"

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

func DeleteWellLogDictCrvCls(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_dict_crv_cls dto.Well_log_dict_crv_cls
	well_log_dict_crv_cls.Curve_class_id = id

	stmt, err := db.Prepare("delete from well_log_dict_crv_cls where curve_class_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_dict_crv_cls.Curve_class_id)
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


