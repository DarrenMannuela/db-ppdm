package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisInterpSurface(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_interp_surface")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_interp_surface

	for rows.Next() {
		var seis_interp_surface dto.Seis_interp_surface
		if err := rows.Scan(&seis_interp_surface.Surface_id, &seis_interp_surface.Active_ind, &seis_interp_surface.Conformity_relationship, &seis_interp_surface.Effective_date, &seis_interp_surface.Expiry_date, &seis_interp_surface.Ordinal_seq_no, &seis_interp_surface.Overturned_ind, &seis_interp_surface.Ppdm_guid, &seis_interp_surface.Remark, &seis_interp_surface.Repeat_strat_occur_no, &seis_interp_surface.Repeat_strat_type, &seis_interp_surface.Source, &seis_interp_surface.Strat_name_set_id, &seis_interp_surface.Strat_unit_id, &seis_interp_surface.Surface_name, &seis_interp_surface.Row_changed_by, &seis_interp_surface.Row_changed_date, &seis_interp_surface.Row_created_by, &seis_interp_surface.Row_created_date, &seis_interp_surface.Row_effective_date, &seis_interp_surface.Row_expiry_date, &seis_interp_surface.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_interp_surface to result
		result = append(result, seis_interp_surface)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisInterpSurface(c *fiber.Ctx) error {
	var seis_interp_surface dto.Seis_interp_surface

	setDefaults(&seis_interp_surface)

	if err := json.Unmarshal(c.Body(), &seis_interp_surface); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_interp_surface values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	seis_interp_surface.Row_created_date = formatDate(seis_interp_surface.Row_created_date)
	seis_interp_surface.Row_changed_date = formatDate(seis_interp_surface.Row_changed_date)
	seis_interp_surface.Effective_date = formatDateString(seis_interp_surface.Effective_date)
	seis_interp_surface.Expiry_date = formatDateString(seis_interp_surface.Expiry_date)
	seis_interp_surface.Row_effective_date = formatDateString(seis_interp_surface.Row_effective_date)
	seis_interp_surface.Row_expiry_date = formatDateString(seis_interp_surface.Row_expiry_date)






	rows, err := stmt.Exec(seis_interp_surface.Surface_id, seis_interp_surface.Active_ind, seis_interp_surface.Conformity_relationship, seis_interp_surface.Effective_date, seis_interp_surface.Expiry_date, seis_interp_surface.Ordinal_seq_no, seis_interp_surface.Overturned_ind, seis_interp_surface.Ppdm_guid, seis_interp_surface.Remark, seis_interp_surface.Repeat_strat_occur_no, seis_interp_surface.Repeat_strat_type, seis_interp_surface.Source, seis_interp_surface.Strat_name_set_id, seis_interp_surface.Strat_unit_id, seis_interp_surface.Surface_name, seis_interp_surface.Row_changed_by, seis_interp_surface.Row_changed_date, seis_interp_surface.Row_created_by, seis_interp_surface.Row_created_date, seis_interp_surface.Row_effective_date, seis_interp_surface.Row_expiry_date, seis_interp_surface.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisInterpSurface(c *fiber.Ctx) error {
	var seis_interp_surface dto.Seis_interp_surface

	setDefaults(&seis_interp_surface)

	if err := json.Unmarshal(c.Body(), &seis_interp_surface); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_interp_surface.Surface_id = id

    if seis_interp_surface.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_interp_surface set active_ind = :1, conformity_relationship = :2, effective_date = :3, expiry_date = :4, ordinal_seq_no = :5, overturned_ind = :6, ppdm_guid = :7, remark = :8, repeat_strat_occur_no = :9, repeat_strat_type = :10, source = :11, strat_name_set_id = :12, strat_unit_id = :13, surface_name = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where surface_id = :22")
	if err != nil {
		return err
	}

	seis_interp_surface.Row_changed_date = formatDate(seis_interp_surface.Row_changed_date)
	seis_interp_surface.Effective_date = formatDateString(seis_interp_surface.Effective_date)
	seis_interp_surface.Expiry_date = formatDateString(seis_interp_surface.Expiry_date)
	seis_interp_surface.Row_effective_date = formatDateString(seis_interp_surface.Row_effective_date)
	seis_interp_surface.Row_expiry_date = formatDateString(seis_interp_surface.Row_expiry_date)






	rows, err := stmt.Exec(seis_interp_surface.Active_ind, seis_interp_surface.Conformity_relationship, seis_interp_surface.Effective_date, seis_interp_surface.Expiry_date, seis_interp_surface.Ordinal_seq_no, seis_interp_surface.Overturned_ind, seis_interp_surface.Ppdm_guid, seis_interp_surface.Remark, seis_interp_surface.Repeat_strat_occur_no, seis_interp_surface.Repeat_strat_type, seis_interp_surface.Source, seis_interp_surface.Strat_name_set_id, seis_interp_surface.Strat_unit_id, seis_interp_surface.Surface_name, seis_interp_surface.Row_changed_by, seis_interp_surface.Row_changed_date, seis_interp_surface.Row_created_by, seis_interp_surface.Row_effective_date, seis_interp_surface.Row_expiry_date, seis_interp_surface.Row_quality, seis_interp_surface.Surface_id)
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

func PatchSeisInterpSurface(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_interp_surface set "
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
	query += " where surface_id = :id"

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

func DeleteSeisInterpSurface(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_interp_surface dto.Seis_interp_surface
	seis_interp_surface.Surface_id = id

	stmt, err := db.Prepare("delete from seis_interp_surface where surface_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_interp_surface.Surface_id)
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


