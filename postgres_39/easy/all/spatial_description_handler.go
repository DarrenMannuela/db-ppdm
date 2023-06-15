package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpatialDescription(c *fiber.Ctx) error {
	rows, err := db.Query("select * from spatial_description")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Spatial_description

	for rows.Next() {
		var spatial_description dto.Spatial_description
		if err := rows.Scan(&spatial_description.Spatial_description_id, &spatial_description.Spatial_obs_no, &spatial_description.Active_ind, &spatial_description.Carter_ind, &spatial_description.Congress_ind, &spatial_description.Coord_acquisition_id, &spatial_description.Coord_system_id, &spatial_description.Dls_ind, &spatial_description.Effective_date, &spatial_description.Expiry_date, &spatial_description.Fps_ind, &spatial_description.Geodetic_ind, &spatial_description.Gross_size, &spatial_description.Gross_size_ouom, &spatial_description.Inactivation_date, &spatial_description.Libya_ind, &spatial_description.Line_ind, &spatial_description.Line_version_ind, &spatial_description.Local_coord_system_id, &spatial_description.Max_latitude, &spatial_description.Max_longitude, &spatial_description.Mineral_zone_ind, &spatial_description.Min_latitude, &spatial_description.Min_longitude, &spatial_description.Ne_loc_ind, &spatial_description.North_sea_ind, &spatial_description.Nts_ind, &spatial_description.Offshore_ind, &spatial_description.Ohio_ind, &spatial_description.Pbl_ind, &spatial_description.Point_ind, &spatial_description.Point_version_ind, &spatial_description.Polygon_ind, &spatial_description.Polygon_version_ind, &spatial_description.Ppdm_guid, &spatial_description.Reference_num, &spatial_description.Remark, &spatial_description.Source, &spatial_description.Spatial_desc_text_ind, &spatial_description.Spatial_desc_type, &spatial_description.Texas_ind, &spatial_description.Row_changed_by, &spatial_description.Row_changed_date, &spatial_description.Row_created_by, &spatial_description.Row_created_date, &spatial_description.Row_effective_date, &spatial_description.Row_expiry_date, &spatial_description.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append spatial_description to result
		result = append(result, spatial_description)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpatialDescription(c *fiber.Ctx) error {
	var spatial_description dto.Spatial_description

	setDefaults(&spatial_description)

	if err := json.Unmarshal(c.Body(), &spatial_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into spatial_description values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48)")
	if err != nil {
		return err
	}
	spatial_description.Row_created_date = formatDate(spatial_description.Row_created_date)
	spatial_description.Row_changed_date = formatDate(spatial_description.Row_changed_date)
	spatial_description.Effective_date = formatDateString(spatial_description.Effective_date)
	spatial_description.Expiry_date = formatDateString(spatial_description.Expiry_date)
	spatial_description.Inactivation_date = formatDateString(spatial_description.Inactivation_date)
	spatial_description.Row_effective_date = formatDateString(spatial_description.Row_effective_date)
	spatial_description.Row_expiry_date = formatDateString(spatial_description.Row_expiry_date)






	rows, err := stmt.Exec(spatial_description.Spatial_description_id, spatial_description.Spatial_obs_no, spatial_description.Active_ind, spatial_description.Carter_ind, spatial_description.Congress_ind, spatial_description.Coord_acquisition_id, spatial_description.Coord_system_id, spatial_description.Dls_ind, spatial_description.Effective_date, spatial_description.Expiry_date, spatial_description.Fps_ind, spatial_description.Geodetic_ind, spatial_description.Gross_size, spatial_description.Gross_size_ouom, spatial_description.Inactivation_date, spatial_description.Libya_ind, spatial_description.Line_ind, spatial_description.Line_version_ind, spatial_description.Local_coord_system_id, spatial_description.Max_latitude, spatial_description.Max_longitude, spatial_description.Mineral_zone_ind, spatial_description.Min_latitude, spatial_description.Min_longitude, spatial_description.Ne_loc_ind, spatial_description.North_sea_ind, spatial_description.Nts_ind, spatial_description.Offshore_ind, spatial_description.Ohio_ind, spatial_description.Pbl_ind, spatial_description.Point_ind, spatial_description.Point_version_ind, spatial_description.Polygon_ind, spatial_description.Polygon_version_ind, spatial_description.Ppdm_guid, spatial_description.Reference_num, spatial_description.Remark, spatial_description.Source, spatial_description.Spatial_desc_text_ind, spatial_description.Spatial_desc_type, spatial_description.Texas_ind, spatial_description.Row_changed_by, spatial_description.Row_changed_date, spatial_description.Row_created_by, spatial_description.Row_created_date, spatial_description.Row_effective_date, spatial_description.Row_expiry_date, spatial_description.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpatialDescription(c *fiber.Ctx) error {
	var spatial_description dto.Spatial_description

	setDefaults(&spatial_description)

	if err := json.Unmarshal(c.Body(), &spatial_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	spatial_description.Spatial_description_id = id

    if spatial_description.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update spatial_description set spatial_obs_no = :1, active_ind = :2, carter_ind = :3, congress_ind = :4, coord_acquisition_id = :5, coord_system_id = :6, dls_ind = :7, effective_date = :8, expiry_date = :9, fps_ind = :10, geodetic_ind = :11, gross_size = :12, gross_size_ouom = :13, inactivation_date = :14, libya_ind = :15, line_ind = :16, line_version_ind = :17, local_coord_system_id = :18, max_latitude = :19, max_longitude = :20, mineral_zone_ind = :21, min_latitude = :22, min_longitude = :23, ne_loc_ind = :24, north_sea_ind = :25, nts_ind = :26, offshore_ind = :27, ohio_ind = :28, pbl_ind = :29, point_ind = :30, point_version_ind = :31, polygon_ind = :32, polygon_version_ind = :33, ppdm_guid = :34, reference_num = :35, remark = :36, source = :37, spatial_desc_text_ind = :38, spatial_desc_type = :39, texas_ind = :40, row_changed_by = :41, row_changed_date = :42, row_created_by = :43, row_effective_date = :44, row_expiry_date = :45, row_quality = :46 where spatial_description_id = :48")
	if err != nil {
		return err
	}

	spatial_description.Row_changed_date = formatDate(spatial_description.Row_changed_date)
	spatial_description.Effective_date = formatDateString(spatial_description.Effective_date)
	spatial_description.Expiry_date = formatDateString(spatial_description.Expiry_date)
	spatial_description.Inactivation_date = formatDateString(spatial_description.Inactivation_date)
	spatial_description.Row_effective_date = formatDateString(spatial_description.Row_effective_date)
	spatial_description.Row_expiry_date = formatDateString(spatial_description.Row_expiry_date)






	rows, err := stmt.Exec(spatial_description.Spatial_obs_no, spatial_description.Active_ind, spatial_description.Carter_ind, spatial_description.Congress_ind, spatial_description.Coord_acquisition_id, spatial_description.Coord_system_id, spatial_description.Dls_ind, spatial_description.Effective_date, spatial_description.Expiry_date, spatial_description.Fps_ind, spatial_description.Geodetic_ind, spatial_description.Gross_size, spatial_description.Gross_size_ouom, spatial_description.Inactivation_date, spatial_description.Libya_ind, spatial_description.Line_ind, spatial_description.Line_version_ind, spatial_description.Local_coord_system_id, spatial_description.Max_latitude, spatial_description.Max_longitude, spatial_description.Mineral_zone_ind, spatial_description.Min_latitude, spatial_description.Min_longitude, spatial_description.Ne_loc_ind, spatial_description.North_sea_ind, spatial_description.Nts_ind, spatial_description.Offshore_ind, spatial_description.Ohio_ind, spatial_description.Pbl_ind, spatial_description.Point_ind, spatial_description.Point_version_ind, spatial_description.Polygon_ind, spatial_description.Polygon_version_ind, spatial_description.Ppdm_guid, spatial_description.Reference_num, spatial_description.Remark, spatial_description.Source, spatial_description.Spatial_desc_text_ind, spatial_description.Spatial_desc_type, spatial_description.Texas_ind, spatial_description.Row_changed_by, spatial_description.Row_changed_date, spatial_description.Row_created_by, spatial_description.Row_effective_date, spatial_description.Row_expiry_date, spatial_description.Row_quality, spatial_description.Spatial_description_id)
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

func PatchSpatialDescription(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update spatial_description set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "inactivation_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where spatial_description_id = :id"

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

func DeleteSpatialDescription(c *fiber.Ctx) error {
	id := c.Params("id")
	var spatial_description dto.Spatial_description
	spatial_description.Spatial_description_id = id

	stmt, err := db.Prepare("delete from spatial_description where spatial_description_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(spatial_description.Spatial_description_id)
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


