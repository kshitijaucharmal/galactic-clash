varying mediump vec2 var_texcoord0;

uniform lowp sampler2D texture_sampler;
uniform lowp vec4 tint;

const float offset = 1.0 / 500.0;
void main()
{
	vec4 col = texture2D(texture_sampler, var_texcoord0.xy);
	lowp vec4 tint_pm = vec4(tint.xyz * tint.w, tint.w);
	if (col.a > 0.5){
		gl_FragColor = col * tint_pm;
	}
	else {
		float a = texture2D(texture_sampler, vec2(var_texcoord0.x + offset, var_texcoord0.y)).a +
		texture2D(texture_sampler, vec2(var_texcoord0.x, var_texcoord0.y - offset)).a +
		texture2D(texture_sampler, vec2(var_texcoord0.x - offset, var_texcoord0.y)).a +
		texture2D(texture_sampler, vec2(var_texcoord0.x, var_texcoord0.y + offset)).a;
		if (col.a < 1.0 && a > 0.0)
		gl_FragColor = vec4(0.0, 0.0, 0.0, 0.8) * tint_pm;
		else
		gl_FragColor = col * tint_pm;
	}
}