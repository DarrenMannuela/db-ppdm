package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisBinOutline(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_bin_outline")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_bin_outline

	for rows.Next() {
		var seis_bin_outline dto.Seis_bin_outline
		if err := rows.Scan(&seis_bin_outline.Seis_set_subtype, &seis_bin_outline.Seis_set_id, &seis_bin_outline.Bin_grid_id, &seis_bin_outline.Bin_grid_source, &seis_bin_outline.Outline_seq_no, &seis_bin_outline.Active_ind, &seis_bin_outline.Bin_outline_type, &seis_bin_outline.Bin_point_id, &seis_bin_outline.Coord_acquisition_id, &seis_bin_outline.Coord_system_id, &seis_bin_outline.Easting, &seis_bin_outline.Effective_date, &seis_bin_outline.Expiry_date, &seis_bin_outline.Local_coord_system_id, &seis_bin_outline.Nline_no, &seis_bin_outline.Northing, &seis_bin_outline.Ppdm_guid, &seis_bin_outline.Remark, &seis_bin_outline.Xline_no, &seis_bin_outline.Row_changed_by, &seis_bin_outline.Row_changed_date, &seis_bin_outline.Row_created_by, &seis_bin_outline.Row_created_date, &seis_bin_outline.Row_effective_date, &seis_bin_outline.Row_expiry_date, &seis_bin_outline.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_bin_outline to result
		result = append(result, seis_bin_outline)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisBinOutline(c *fiber.Ctx) error {
	var seis_bin_outline dto.Seis_bin_outline

	setDefaults(&seis_bin_outline)

	if err := json.Unmarshal(c.Body(), &seis_bin_outline); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_bin_outline values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	seis_bin_outline.Row_created_date = formatDate(seis_bin_outline.Row_created_date)
	seis_bin_outline.Row_changed_date = formatDate(seis_bin_outline.Row_changed_date)
	seis_bin_outline.Effective_date = formatDateString(seis_bin_outline.Effective_date)
	seis_bin_outline.Expiry_date = formatDateString(seis_bin_outline.Expiry_date)
	seis_bin_outline.Row_effective_date = formatDateString(seis_bin_outline.Row_effective_date)
	seis_bin_outline.Row_expiry_date = formatDateString(seis_bin_outline.Row_expiry_date)






	rows, err := stmt.Exec(seis_bin_outline.Seis_set_subtype, seis_bin_outline.Seis_set_id, seis_bin_outline.Bin_grid_id, seis_bin_outline.Bin_grid_source, seis_bin_outline.Outline_seq_no, seis_bin_outline.Active_ind, seis_bin_outline.Bin_outline_type, seis_bin_outline.Bin_point_id, seis_bin_outline.Coord_acquisition_id, seis_bin_outline.Coord_system_id, seis_bin_outline.Easting, seis_bin_outline.Effective_date, seis_bin_outline.Expiry_date, seis_bin_outline.Local_coord_system_id, seis_bin_outline.Nline_no, seis_bin_outline.Northing, seis_bin_outline.Ppdm_guid, seis_bin_outline.Remark, seis_bin_outline.Xline_no, seis_bin_outline.Row_changed_by, seis_bin_outline.Row_changed_date, seis_bin_outline.Row_created_by, seis_bin_outline.Row_created_date, seis_bin_outline.Row_effective_date, seis_bin_outline.Row_expiry_date, seis_bin_outline.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisBinOutline(c *fiber.Ctx) error {
	var seis_bin_outline dto.Seis_bin_outline

	setDefaults(&seis_bin_outline)

	if err := json.Unmarshal(c.Body(), &seis_bin_outline); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_bin_outline.Seis_set_subtype = id

    if seis_bin_outline.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_bin_outline set seis_set_id = :1, bin_grid_id = :2, bin_grid_source = :3, outline_seq_no = :4, active_ind = :5, bin_outline_type = :6, bin_point_id = :7, coord_acquisition_id = :8, coord_system_id = :9, easting = :10, effective_date = :11, expiry_date = :12, local_coord_system_id = :13, nline_no = :14, northing = :15, ppdm_guid = :16, remark = :17, xline_no = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where seis_set_subtype = :26")
	if err != nil {
		return err
	}

	seis_bin_outline.Row_changed_date = formatDate(seis_bin_outline.Row_changed_date)
	seis_bin_outline.Effective_date = formatDateString(seis_bin_outline.Effective_date)
	seis_bin_outline.Expiry_date = formatDateString(seis_bin_outline.Expiry_date)
	seis_bin_outline.Row_effective_date = formatDateString(seis_bin_outline.Row_effective_date)
	seis_bin_outline.Row_expiry_date = formatDateString(seis_bin_outline.Row_expiry_date)






	rows, err := stmt.Exec(seis_bin_outline.Seis_set_id, seis_bin_outline.Bin_grid_id, seis_bin_outline.Bin_grid_source, seis_bin_outline.Outline_seq_no, seis_bin_outline.Active_ind, seis_bin_outline.Bin_outline_type, seis_bin_outline.Bin_point_id, seis_bin_outline.Coord_acquisition_id, seis_bin_outline.Coord_system_id, seis_bin_outline.Easting, seis_bin_outline.Effective_date, seis_bin_outline.Expiry_date, seis_bin_outline.Local_coord_system_id, seis_bin_outline.Nline_no, seis_bin_outline.Northing, seis_bin_outline.Ppdm_guid, seis_bin_outline.Remark, seis_bin_outline.Xline_no, seis_bin_outline.Row_changed_by, seis_bin_outline.Row_changed_date, seis_bin_outline.Row_created_by, seis_bin_outline.Row_effective_date, seis_bin_outline.Row_expiry_date, seis_bin_outline.Row_quality, seis_bin_outline.Seis_set_subtype)
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

func PatchSeisBinOutline(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_bin_outline set "
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

func DeleteSeisBinOutline(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_bin_outline dto.Seis_bin_outline
	seis_bin_outline.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_bin_outline where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_bin_outline.Seis_set_subtype)
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


