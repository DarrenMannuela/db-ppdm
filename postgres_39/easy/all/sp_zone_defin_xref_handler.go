package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpZoneDefinXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_zone_defin_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_zone_defin_xref

	for rows.Next() {
		var sp_zone_defin_xref dto.Sp_zone_defin_xref
		if err := rows.Scan(&sp_zone_defin_xref.Zone_definition_id_1, &sp_zone_defin_xref.Zone_definition_id_2, &sp_zone_defin_xref.Active_ind, &sp_zone_defin_xref.Description, &sp_zone_defin_xref.Effective_date, &sp_zone_defin_xref.Expiry_date, &sp_zone_defin_xref.Ppdm_guid, &sp_zone_defin_xref.Remark, &sp_zone_defin_xref.Responsible_ba_id, &sp_zone_defin_xref.Source, &sp_zone_defin_xref.Xref_reason, &sp_zone_defin_xref.Row_changed_by, &sp_zone_defin_xref.Row_changed_date, &sp_zone_defin_xref.Row_created_by, &sp_zone_defin_xref.Row_created_date, &sp_zone_defin_xref.Row_effective_date, &sp_zone_defin_xref.Row_expiry_date, &sp_zone_defin_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_zone_defin_xref to result
		result = append(result, sp_zone_defin_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpZoneDefinXref(c *fiber.Ctx) error {
	var sp_zone_defin_xref dto.Sp_zone_defin_xref

	setDefaults(&sp_zone_defin_xref)

	if err := json.Unmarshal(c.Body(), &sp_zone_defin_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_zone_defin_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	sp_zone_defin_xref.Row_created_date = formatDate(sp_zone_defin_xref.Row_created_date)
	sp_zone_defin_xref.Row_changed_date = formatDate(sp_zone_defin_xref.Row_changed_date)
	sp_zone_defin_xref.Effective_date = formatDateString(sp_zone_defin_xref.Effective_date)
	sp_zone_defin_xref.Expiry_date = formatDateString(sp_zone_defin_xref.Expiry_date)
	sp_zone_defin_xref.Row_effective_date = formatDateString(sp_zone_defin_xref.Row_effective_date)
	sp_zone_defin_xref.Row_expiry_date = formatDateString(sp_zone_defin_xref.Row_expiry_date)






	rows, err := stmt.Exec(sp_zone_defin_xref.Zone_definition_id_1, sp_zone_defin_xref.Zone_definition_id_2, sp_zone_defin_xref.Active_ind, sp_zone_defin_xref.Description, sp_zone_defin_xref.Effective_date, sp_zone_defin_xref.Expiry_date, sp_zone_defin_xref.Ppdm_guid, sp_zone_defin_xref.Remark, sp_zone_defin_xref.Responsible_ba_id, sp_zone_defin_xref.Source, sp_zone_defin_xref.Xref_reason, sp_zone_defin_xref.Row_changed_by, sp_zone_defin_xref.Row_changed_date, sp_zone_defin_xref.Row_created_by, sp_zone_defin_xref.Row_created_date, sp_zone_defin_xref.Row_effective_date, sp_zone_defin_xref.Row_expiry_date, sp_zone_defin_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpZoneDefinXref(c *fiber.Ctx) error {
	var sp_zone_defin_xref dto.Sp_zone_defin_xref

	setDefaults(&sp_zone_defin_xref)

	if err := json.Unmarshal(c.Body(), &sp_zone_defin_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_zone_defin_xref.Zone_definition_id_1 = id

    if sp_zone_defin_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_zone_defin_xref set zone_definition_id_2 = :1, active_ind = :2, description = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, responsible_ba_id = :8, source = :9, xref_reason = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where zone_definition_id_1 = :18")
	if err != nil {
		return err
	}

	sp_zone_defin_xref.Row_changed_date = formatDate(sp_zone_defin_xref.Row_changed_date)
	sp_zone_defin_xref.Effective_date = formatDateString(sp_zone_defin_xref.Effective_date)
	sp_zone_defin_xref.Expiry_date = formatDateString(sp_zone_defin_xref.Expiry_date)
	sp_zone_defin_xref.Row_effective_date = formatDateString(sp_zone_defin_xref.Row_effective_date)
	sp_zone_defin_xref.Row_expiry_date = formatDateString(sp_zone_defin_xref.Row_expiry_date)






	rows, err := stmt.Exec(sp_zone_defin_xref.Zone_definition_id_2, sp_zone_defin_xref.Active_ind, sp_zone_defin_xref.Description, sp_zone_defin_xref.Effective_date, sp_zone_defin_xref.Expiry_date, sp_zone_defin_xref.Ppdm_guid, sp_zone_defin_xref.Remark, sp_zone_defin_xref.Responsible_ba_id, sp_zone_defin_xref.Source, sp_zone_defin_xref.Xref_reason, sp_zone_defin_xref.Row_changed_by, sp_zone_defin_xref.Row_changed_date, sp_zone_defin_xref.Row_created_by, sp_zone_defin_xref.Row_effective_date, sp_zone_defin_xref.Row_expiry_date, sp_zone_defin_xref.Row_quality, sp_zone_defin_xref.Zone_definition_id_1)
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

func PatchSpZoneDefinXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_zone_defin_xref set "
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
	query += " where zone_definition_id_1 = :id"

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

func DeleteSpZoneDefinXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_zone_defin_xref dto.Sp_zone_defin_xref
	sp_zone_defin_xref.Zone_definition_id_1 = id

	stmt, err := db.Prepare("delete from sp_zone_defin_xref where zone_definition_id_1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_zone_defin_xref.Zone_definition_id_1)
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


