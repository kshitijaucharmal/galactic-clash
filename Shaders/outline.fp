varying mediump vec2 var_texcoord0;

uniform lowp sampler2D texture_sampler;

const float offset = 1.0 / 250.0;
void main()
{
	vec4 col = texture2D(texture_sampler, var_texcoord0.xy);
	if (col.a > 0.5)
	gl_FragColor = col;
	else {
		float a = texture2D(texture_sampler, vec2(var_texcoord0.x + offset, var_texcoord0.y)).a +
		texture2D(texture_sampler, vec2(var_texcoord0.x, var_texcoord0.y - offset)).a +
		texture2D(texture_sampler, vec2(var_texcoord0.x - offset, var_texcoord0.y)).a +
		texture2D(texture_sampler, vec2(var_texcoord0.x, var_texcoord0.y + offset)).a;
		if (col.a < 1.0 && a > 0.0)
		gl_FragColor = vec4(0.0, 0.0, 0.0, 0.8);
		else
		gl_FragColor = col;
	}
}