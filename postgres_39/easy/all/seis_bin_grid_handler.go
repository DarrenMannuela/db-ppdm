package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisBinGrid(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_bin_grid")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_bin_grid

	for rows.Next() {
		var seis_bin_grid dto.Seis_bin_grid
		if err := rows.Scan(&seis_bin_grid.Seis_set_subtype, &seis_bin_grid.Seis_set_id, &seis_bin_grid.Bin_grid_id, &seis_bin_grid.Source, &seis_bin_grid.Active_ind, &seis_bin_grid.Angle_rotation, &seis_bin_grid.Bin_grid_type, &seis_bin_grid.Bin_grid_version, &seis_bin_grid.Bin_method, &seis_bin_grid.Bin_point_ouom, &seis_bin_grid.Coord_acquisition_id, &seis_bin_grid.Coord_system_id, &seis_bin_grid.Corner1_lat, &seis_bin_grid.Corner1_long, &seis_bin_grid.Corner2_lat, &seis_bin_grid.Corner2_long, &seis_bin_grid.Corner3_lat, &seis_bin_grid.Corner3_long, &seis_bin_grid.Effective_date, &seis_bin_grid.Expiry_date, &seis_bin_grid.Fold_coverage, &seis_bin_grid.Local_coord_system_id, &seis_bin_grid.Nline_azimuth, &seis_bin_grid.Nline_count, &seis_bin_grid.Nline_max_no, &seis_bin_grid.Nline_min_no, &seis_bin_grid.Nline_spacing, &seis_bin_grid.North_type, &seis_bin_grid.Point_origin_easting, &seis_bin_grid.Point_origin_latitude, &seis_bin_grid.Point_origin_longitude, &seis_bin_grid.Point_origin_northing, &seis_bin_grid.Ppdm_guid, &seis_bin_grid.Remark, &seis_bin_grid.Xline_azimuth, &seis_bin_grid.Xline_count, &seis_bin_grid.Xline_max_no, &seis_bin_grid.Xline_min_no, &seis_bin_grid.Xline_spacing, &seis_bin_grid.Row_changed_by, &seis_bin_grid.Row_changed_date, &seis_bin_grid.Row_created_by, &seis_bin_grid.Row_created_date, &seis_bin_grid.Row_effective_date, &seis_bin_grid.Row_expiry_date, &seis_bin_grid.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_bin_grid to result
		result = append(result, seis_bin_grid)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisBinGrid(c *fiber.Ctx) error {
	var seis_bin_grid dto.Seis_bin_grid

	setDefaults(&seis_bin_grid)

	if err := json.Unmarshal(c.Body(), &seis_bin_grid); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_bin_grid values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46)")
	if err != nil {
		return err
	}
	seis_bin_grid.Row_created_date = formatDate(seis_bin_grid.Row_created_date)
	seis_bin_grid.Row_changed_date = formatDate(seis_bin_grid.Row_changed_date)
	seis_bin_grid.Effective_date = formatDateString(seis_bin_grid.Effective_date)
	seis_bin_grid.Expiry_date = formatDateString(seis_bin_grid.Expiry_date)
	seis_bin_grid.Row_effective_date = formatDateString(seis_bin_grid.Row_effective_date)
	seis_bin_grid.Row_expiry_date = formatDateString(seis_bin_grid.Row_expiry_date)






	rows, err := stmt.Exec(seis_bin_grid.Seis_set_subtype, seis_bin_grid.Seis_set_id, seis_bin_grid.Bin_grid_id, seis_bin_grid.Source, seis_bin_grid.Active_ind, seis_bin_grid.Angle_rotation, seis_bin_grid.Bin_grid_type, seis_bin_grid.Bin_grid_version, seis_bin_grid.Bin_method, seis_bin_grid.Bin_point_ouom, seis_bin_grid.Coord_acquisition_id, seis_bin_grid.Coord_system_id, seis_bin_grid.Corner1_lat, seis_bin_grid.Corner1_long, seis_bin_grid.Corner2_lat, seis_bin_grid.Corner2_long, seis_bin_grid.Corner3_lat, seis_bin_grid.Corner3_long, seis_bin_grid.Effective_date, seis_bin_grid.Expiry_date, seis_bin_grid.Fold_coverage, seis_bin_grid.Local_coord_system_id, seis_bin_grid.Nline_azimuth, seis_bin_grid.Nline_count, seis_bin_grid.Nline_max_no, seis_bin_grid.Nline_min_no, seis_bin_grid.Nline_spacing, seis_bin_grid.North_type, seis_bin_grid.Point_origin_easting, seis_bin_grid.Point_origin_latitude, seis_bin_grid.Point_origin_longitude, seis_bin_grid.Point_origin_northing, seis_bin_grid.Ppdm_guid, seis_bin_grid.Remark, seis_bin_grid.Xline_azimuth, seis_bin_grid.Xline_count, seis_bin_grid.Xline_max_no, seis_bin_grid.Xline_min_no, seis_bin_grid.Xline_spacing, seis_bin_grid.Row_changed_by, seis_bin_grid.Row_changed_date, seis_bin_grid.Row_created_by, seis_bin_grid.Row_created_date, seis_bin_grid.Row_effective_date, seis_bin_grid.Row_expiry_date, seis_bin_grid.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisBinGrid(c *fiber.Ctx) error {
	var seis_bin_grid dto.Seis_bin_grid

	setDefaults(&seis_bin_grid)

	if err := json.Unmarshal(c.Body(), &seis_bin_grid); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_bin_grid.Seis_set_subtype = id

    if seis_bin_grid.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_bin_grid set seis_set_id = :1, bin_grid_id = :2, source = :3, active_ind = :4, angle_rotation = :5, bin_grid_type = :6, bin_grid_version = :7, bin_method = :8, bin_point_ouom = :9, coord_acquisition_id = :10, coord_system_id = :11, corner1_lat = :12, corner1_long = :13, corner2_lat = :14, corner2_long = :15, corner3_lat = :16, corner3_long = :17, effective_date = :18, expiry_date = :19, fold_coverage = :20, local_coord_system_id = :21, nline_azimuth = :22, nline_count = :23, nline_max_no = :24, nline_min_no = :25, nline_spacing = :26, north_type = :27, point_origin_easting = :28, point_origin_latitude = :29, point_origin_longitude = :30, point_origin_northing = :31, ppdm_guid = :32, remark = :33, xline_azimuth = :34, xline_count = :35, xline_max_no = :36, xline_min_no = :37, xline_spacing = :38, row_changed_by = :39, row_changed_date = :40, row_created_by = :41, row_effective_date = :42, row_expiry_date = :43, row_quality = :44 where seis_set_subtype = :46")
	if err != nil {
		return err
	}

	seis_bin_grid.Row_changed_date = formatDate(seis_bin_grid.Row_changed_date)
	seis_bin_grid.Effective_date = formatDateString(seis_bin_grid.Effective_date)
	seis_bin_grid.Expiry_date = formatDateString(seis_bin_grid.Expiry_date)
	seis_bin_grid.Row_effective_date = formatDateString(seis_bin_grid.Row_effective_date)
	seis_bin_grid.Row_expiry_date = formatDateString(seis_bin_grid.Row_expiry_date)






	rows, err := stmt.Exec(seis_bin_grid.Seis_set_id, seis_bin_grid.Bin_grid_id, seis_bin_grid.Source, seis_bin_grid.Active_ind, seis_bin_grid.Angle_rotation, seis_bin_grid.Bin_grid_type, seis_bin_grid.Bin_grid_version, seis_bin_grid.Bin_method, seis_bin_grid.Bin_point_ouom, seis_bin_grid.Coord_acquisition_id, seis_bin_grid.Coord_system_id, seis_bin_grid.Corner1_lat, seis_bin_grid.Corner1_long, seis_bin_grid.Corner2_lat, seis_bin_grid.Corner2_long, seis_bin_grid.Corner3_lat, seis_bin_grid.Corner3_long, seis_bin_grid.Effective_date, seis_bin_grid.Expiry_date, seis_bin_grid.Fold_coverage, seis_bin_grid.Local_coord_system_id, seis_bin_grid.Nline_azimuth, seis_bin_grid.Nline_count, seis_bin_grid.Nline_max_no, seis_bin_grid.Nline_min_no, seis_bin_grid.Nline_spacing, seis_bin_grid.North_type, seis_bin_grid.Point_origin_easting, seis_bin_grid.Point_origin_latitude, seis_bin_grid.Point_origin_longitude, seis_bin_grid.Point_origin_northing, seis_bin_grid.Ppdm_guid, seis_bin_grid.Remark, seis_bin_grid.Xline_azimuth, seis_bin_grid.Xline_count, seis_bin_grid.Xline_max_no, seis_bin_grid.Xline_min_no, seis_bin_grid.Xline_spacing, seis_bin_grid.Row_changed_by, seis_bin_grid.Row_changed_date, seis_bin_grid.Row_created_by, seis_bin_grid.Row_effective_date, seis_bin_grid.Row_expiry_date, seis_bin_grid.Row_quality, seis_bin_grid.Seis_set_subtype)
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

func PatchSeisBinGrid(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_bin_grid set "
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

func DeleteSeisBinGrid(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_bin_grid dto.Seis_bin_grid
	seis_bin_grid.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_bin_grid where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_bin_grid.Seis_set_subtype)
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


