update schemas
set palette = palette - (json_array_length(palette) - 1);
