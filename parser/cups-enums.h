typedef enum cups_adv_e			/**** AdvanceMedia attribute values ****/
{
  CUPS_ADVANCE_NONE = 0,		/* Never advance the roll */
  CUPS_ADVANCE_FILE = 1,		/* Advance the roll after this file */
  CUPS_ADVANCE_JOB = 2,			/* Advance the roll after this job */
  CUPS_ADVANCE_SET = 3,			/* Advance the roll after this set */
  CUPS_ADVANCE_PAGE = 4			/* Advance the roll after this page */
} cups_adv_t;

typedef enum cups_bool_e		/**** Boolean type ****/
{
  CUPS_FALSE = 0,			/* Logical false */
  CUPS_TRUE = 1				/* Logical true */
} cups_bool_t;

typedef enum cups_cspace_e		/**** cupsColorSpace attribute values ****/
{
  CUPS_CSPACE_W = 0,			/* Luminance (DeviceGray, gamma 2.2 by default) */
  CUPS_CSPACE_RGB = 1,			/* Red, green, blue (DeviceRGB, sRGB by default) */
  CUPS_CSPACE_RGBA = 2,			/* Red, green, blue, alpha (DeviceRGB, sRGB by default) */
  CUPS_CSPACE_K = 3,			/* Black (DeviceK) */
  CUPS_CSPACE_CMY = 4,			/* Cyan, magenta, yellow (DeviceCMY) */
  CUPS_CSPACE_YMC = 5,			/* Yellow, magenta, cyan @deprecated@ */
  CUPS_CSPACE_CMYK = 6,			/* Cyan, magenta, yellow, black (DeviceCMYK) */
  CUPS_CSPACE_YMCK = 7,			/* Yellow, magenta, cyan, black @deprecated@ */
  CUPS_CSPACE_KCMY = 8,			/* Black, cyan, magenta, yellow @deprecated@ */
  CUPS_CSPACE_KCMYcm = 9,		/* Black, cyan, magenta, yellow, light-cyan, light-magenta @deprecated@ */
  CUPS_CSPACE_GMCK = 10,		/* Gold, magenta, yellow, black @deprecated@ */
  CUPS_CSPACE_GMCS = 11,		/* Gold, magenta, yellow, silver @deprecated@ */
  CUPS_CSPACE_WHITE = 12,		/* White ink (as black) @deprecated@ */
  CUPS_CSPACE_GOLD = 13,		/* Gold foil @deprecated@ */
  CUPS_CSPACE_SILVER = 14,		/* Silver foil @deprecated@ */

  CUPS_CSPACE_CIEXYZ = 15,		/* CIE XYZ @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_CIELab = 16,		/* CIE Lab @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_RGBW = 17,		/* Red, green, blue, white (DeviceRGB, sRGB by default) @since CUPS 1.2/macOS 10.5@ */
  CUPS_CSPACE_SW = 18,			/* Luminance (gamma 2.2) @since CUPS 1.4.5@ */
  CUPS_CSPACE_SRGB = 19,		/* Red, green, blue (sRGB) @since CUPS 1.4.5@ */
  CUPS_CSPACE_ADOBERGB = 20,		/* Red, green, blue (Adobe RGB) @since CUPS 1.4.5@ */

  CUPS_CSPACE_ICC1 = 32,		/* ICC-based, 1 color @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICC2 = 33,		/* ICC-based, 2 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICC3 = 34,		/* ICC-based, 3 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICC4 = 35,		/* ICC-based, 4 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICC5 = 36,		/* ICC-based, 5 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICC6 = 37,		/* ICC-based, 6 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICC7 = 38,		/* ICC-based, 7 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICC8 = 39,		/* ICC-based, 8 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICC9 = 40,		/* ICC-based, 9 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICCA = 41,		/* ICC-based, 10 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICCB = 42,		/* ICC-based, 11 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICCC = 43,		/* ICC-based, 12 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICCD = 44,		/* ICC-based, 13 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICCE = 45,		/* ICC-based, 14 colors @since CUPS 1.1.19/macOS 10.3@ */
  CUPS_CSPACE_ICCF = 46,		/* ICC-based, 15 colors @since CUPS 1.1.19/macOS 10.3@ */

  CUPS_CSPACE_DEVICE1 = 48,		/* DeviceN, 1 color @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICE2 = 49,		/* DeviceN, 2 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICE3 = 50,		/* DeviceN, 3 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICE4 = 51,		/* DeviceN, 4 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICE5 = 52,		/* DeviceN, 5 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICE6 = 53,		/* DeviceN, 6 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICE7 = 54,		/* DeviceN, 7 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICE8 = 55,		/* DeviceN, 8 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICE9 = 56,		/* DeviceN, 9 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICEA = 57,		/* DeviceN, 10 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICEB = 58,		/* DeviceN, 11 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICEC = 59,		/* DeviceN, 12 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICED = 60,		/* DeviceN, 13 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICEE = 61,		/* DeviceN, 14 colors @since CUPS 1.4.5@ */
  CUPS_CSPACE_DEVICEF = 62		/* DeviceN, 15 colors @since CUPS 1.4.5@ */
} cups_cspace_t;

typedef enum cups_cut_e			/**** CutMedia attribute values ****/
{
  CUPS_CUT_NONE = 0,			/* Never cut the roll */
  CUPS_CUT_FILE = 1,			/* Cut the roll after this file */
  CUPS_CUT_JOB = 2,			/* Cut the roll after this job */
  CUPS_CUT_SET = 3,			/* Cut the roll after this set */
  CUPS_CUT_PAGE = 4			/* Cut the roll after this page */
} cups_cut_t;

typedef enum cups_edge_e		/**** LeadingEdge attribute values ****/
{
  CUPS_EDGE_TOP = 0,			/* Leading edge is the top of the page */
  CUPS_EDGE_RIGHT = 1,			/* Leading edge is the right of the page */
  CUPS_EDGE_BOTTOM = 2,			/* Leading edge is the bottom of the page */
  CUPS_EDGE_LEFT = 3			/* Leading edge is the left of the page */
} cups_edge_t;

typedef enum cups_jog_e			/**** Jog attribute values ****/
{
  CUPS_JOG_NONE = 0,			/* Never move pages */
  CUPS_JOG_FILE = 1,			/* Move pages after this file */
  CUPS_JOG_JOB = 2,			/* Move pages after this job */
  CUPS_JOG_SET = 3			/* Move pages after this set */
} cups_jog_t;

enum cups_mode_e			/**** cupsRasterOpen modes ****/
{
  CUPS_RASTER_READ = 0,			/* Open stream for reading */
  CUPS_RASTER_WRITE = 1,		/* Open stream for writing */
  CUPS_RASTER_WRITE_COMPRESSED = 2,	/* Open stream for compressed writing @since CUPS 1.3/macOS 10.5@ */
  CUPS_RASTER_WRITE_PWG = 3		/* Open stream for compressed writing in PWG mode @since CUPS 1.5/macOS 10.7@ */
};

typedef enum cups_mode_e cups_mode_t;	/**** cupsRasterOpen modes ****/

typedef enum cups_order_e		/**** cupsColorOrder attribute values ****/
{
  CUPS_ORDER_CHUNKED = 0,		/* CMYK CMYK CMYK ... */
  CUPS_ORDER_BANDED = 1,		/* CCC MMM YYY KKK ... */
  CUPS_ORDER_PLANAR = 2			/* CCC ... MMM ... YYY ... KKK ... */
} cups_order_t;

typedef enum cups_orient_e		/**** Orientation attribute values ****/
{
  CUPS_ORIENT_0 = 0,			/* Don't rotate the page */
  CUPS_ORIENT_90 = 1,			/* Rotate the page counter-clockwise */
  CUPS_ORIENT_180 = 2,			/* Turn the page upside down */
  CUPS_ORIENT_270 = 3			/* Rotate the page clockwise */
} cups_orient_t;