package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisBinOrigin(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_bin_origin")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_bin_origin

	for rows.Next() {
		var seis_bin_origin dto.Seis_bin_origin
		if err := rows.Scan(&seis_bin_origin.Seis_set_subtype, &seis_bin_origin.Seis_set_id, &seis_bin_origin.Bin_grid_id, &seis_bin_origin.Bin_grid_source, &seis_bin_origin.Bin_origin_id, &seis_bin_origin.Active_ind, &seis_bin_origin.Coord_acquisition_id, &seis_bin_origin.Coord_system_id, &seis_bin_origin.Corner1_lat, &seis_bin_origin.Corner1_long, &seis_bin_origin.Corner2_lat, &seis_bin_origin.Corner2_long, &seis_bin_origin.Corner3_lat, &seis_bin_origin.Corner3_long, &seis_bin_origin.Effective_date, &seis_bin_origin.Exclusion_ind, &seis_bin_origin.Expiry_date, &seis_bin_origin.Inclusion_ind, &seis_bin_origin.Input_bin_grid_id, &seis_bin_origin.Input_bin_source, &seis_bin_origin.Input_seis_set_id, &seis_bin_origin.Input_seis_set_subtype, &seis_bin_origin.Local_coord_system_id, &seis_bin_origin.Ppdm_guid, &seis_bin_origin.Remark, &seis_bin_origin.Source, &seis_bin_origin.Row_changed_by, &seis_bin_origin.Row_changed_date, &seis_bin_origin.Row_created_by, &seis_bin_origin.Row_created_date, &seis_bin_origin.Row_effective_date, &seis_bin_origin.Row_expiry_date, &seis_bin_origin.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_bin_origin to result
		result = append(result, seis_bin_origin)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisBinOrigin(c *fiber.Ctx) error {
	var seis_bin_origin dto.Seis_bin_origin

	setDefaults(&seis_bin_origin)

	if err := json.Unmarshal(c.Body(), &seis_bin_origin); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_bin_origin values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	seis_bin_origin.Row_created_date = formatDate(seis_bin_origin.Row_created_date)
	seis_bin_origin.Row_changed_date = formatDate(seis_bin_origin.Row_changed_date)
	seis_bin_origin.Effective_date = formatDateString(seis_bin_origin.Effective_date)
	seis_bin_origin.Expiry_date = formatDateString(seis_bin_origin.Expiry_date)
	seis_bin_origin.Row_effective_date = formatDateString(seis_bin_origin.Row_effective_date)
	seis_bin_origin.Row_expiry_date = formatDateString(seis_bin_origin.Row_expiry_date)






	rows, err := stmt.Exec(seis_bin_origin.Seis_set_subtype, seis_bin_origin.Seis_set_id, seis_bin_origin.Bin_grid_id, seis_bin_origin.Bin_grid_source, seis_bin_origin.Bin_origin_id, seis_bin_origin.Active_ind, seis_bin_origin.Coord_acquisition_id, seis_bin_origin.Coord_system_id, seis_bin_origin.Corner1_lat, seis_bin_origin.Corner1_long, seis_bin_origin.Corner2_lat, seis_bin_origin.Corner2_long, seis_bin_origin.Corner3_lat, seis_bin_origin.Corner3_long, seis_bin_origin.Effective_date, seis_bin_origin.Exclusion_ind, seis_bin_origin.Expiry_date, seis_bin_origin.Inclusion_ind, seis_bin_origin.Input_bin_grid_id, seis_bin_origin.Input_bin_source, seis_bin_origin.Input_seis_set_id, seis_bin_origin.Input_seis_set_subtype, seis_bin_origin.Local_coord_system_id, seis_bin_origin.Ppdm_guid, seis_bin_origin.Remark, seis_bin_origin.Source, seis_bin_origin.Row_changed_by, seis_bin_origin.Row_changed_date, seis_bin_origin.Row_created_by, seis_bin_origin.Row_created_date, seis_bin_origin.Row_effective_date, seis_bin_origin.Row_expiry_date, seis_bin_origin.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisBinOrigin(c *fiber.Ctx) error {
	var seis_bin_origin dto.Seis_bin_origin

	setDefaults(&seis_bin_origin)

	if err := json.Unmarshal(c.Body(), &seis_bin_origin); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_bin_origin.Seis_set_subtype = id

    if seis_bin_origin.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_bin_origin set seis_set_id = :1, bin_grid_id = :2, bin_grid_source = :3, bin_origin_id = :4, active_ind = :5, coord_acquisition_id = :6, coord_system_id = :7, corner1_lat = :8, corner1_long = :9, corner2_lat = :10, corner2_long = :11, corner3_lat = :12, corner3_long = :13, effective_date = :14, exclusion_ind = :15, expiry_date = :16, inclusion_ind = :17, input_bin_grid_id = :18, input_bin_source = :19, input_seis_set_id = :20, input_seis_set_subtype = :21, local_coord_system_id = :22, ppdm_guid = :23, remark = :24, source = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where seis_set_subtype = :33")
	if err != nil {
		return err
	}

	seis_bin_origin.Row_changed_date = formatDate(seis_bin_origin.Row_changed_date)
	seis_bin_origin.Effective_date = formatDateString(seis_bin_origin.Effective_date)
	seis_bin_origin.Expiry_date = formatDateString(seis_bin_origin.Expiry_date)
	seis_bin_origin.Row_effective_date = formatDateString(seis_bin_origin.Row_effective_date)
	seis_bin_origin.Row_expiry_date = formatDateString(seis_bin_origin.Row_expiry_date)






	rows, err := stmt.Exec(seis_bin_origin.Seis_set_id, seis_bin_origin.Bin_grid_id, seis_bin_origin.Bin_grid_source, seis_bin_origin.Bin_origin_id, seis_bin_origin.Active_ind, seis_bin_origin.Coord_acquisition_id, seis_bin_origin.Coord_system_id, seis_bin_origin.Corner1_lat, seis_bin_origin.Corner1_long, seis_bin_origin.Corner2_lat, seis_bin_origin.Corner2_long, seis_bin_origin.Corner3_lat, seis_bin_origin.Corner3_long, seis_bin_origin.Effective_date, seis_bin_origin.Exclusion_ind, seis_bin_origin.Expiry_date, seis_bin_origin.Inclusion_ind, seis_bin_origin.Input_bin_grid_id, seis_bin_origin.Input_bin_source, seis_bin_origin.Input_seis_set_id, seis_bin_origin.Input_seis_set_subtype, seis_bin_origin.Local_coord_system_id, seis_bin_origin.Ppdm_guid, seis_bin_origin.Remark, seis_bin_origin.Source, seis_bin_origin.Row_changed_by, seis_bin_origin.Row_changed_date, seis_bin_origin.Row_created_by, seis_bin_origin.Row_effective_date, seis_bin_origin.Row_expiry_date, seis_bin_origin.Row_quality, seis_bin_origin.Seis_set_subtype)
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

func PatchSeisBinOrigin(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_bin_origin set "
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

func DeleteSeisBinOrigin(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_bin_origin dto.Seis_bin_origin
	seis_bin_origin.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_bin_origin where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_bin_origin.Seis_set_subtype)
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


