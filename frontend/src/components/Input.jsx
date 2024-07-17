const InputField = ({ className, type, placeholder, value, onChange }) => {
    const classes = `input-field border border-gray-300 rounded-md px-4 py-2 focus:outline-none focus:border-color-1 ${className || ""}`;
  
    return (
      <input
        type={type || "text"}
        className={classes}
        placeholder={placeholder}
        value={value}
        onChange={onChange}
      />
    );
  };
  
  export default InputField;
  